#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, re, json
from shutil import copy, copytree, rmtree
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Dict, Any
from itertools import groupby

RE_HINT = re.compile(r"\s*@([a-z]+)(?::?\s*(.*))?$")
RE_SINGLE_SPACE = re.compile(r"^\s(?:[^\s]*|$)")


class SearchIndexFile(object):
    def __init__(self, tag: str, title: str | None = None):
        self.tag = tag
        self.title = title or ""

    def __hash__(self):
        return hash(self.tag)

    def __eq__(self, value):
        if not isinstance(value, SearchIndexFile):
            return False
        return self.tag == value.tag


class SearchIndexEntry:
    def __init__(self, count: int = 1, comment: str | None = None):
        self.count = count
        self.comment = comment

    def update(self, comment: str | None = None):
        self.count += 1
        if comment:
            comment = re.sub(r"\s+", " ", comment).strip()
        if comment and self.comment:
            self.comment = self.comment + comment
        else:
            self.comment = comment or self.comment


class SearchIndex:
    def __init__(self):
        self.data: dict[SearchIndexFile, dict[str, SearchIndexEntry]] = {}

    def add(
        self,
        filename: str,
        word: str | List[str],
        comment: str | None = None,
        title: str | None = None,
    ):
        if isinstance(word, str) and word.startswith("google.protobuf."):
            return

        tag, data = next(
            ((k, v) for k, v in self.data.items() if k.tag == filename), (None, None)
        )
        if tag is None:
            data = self.data[SearchIndexFile(filename, title)] = {}
        else:
            tag.title = title or tag.title

        if isinstance(word, list):
            for w in word:
                self.add(filename, w)
            return
        if word not in data:
            data[word] = SearchIndexEntry(count=1, comment=comment)
        else:
            data[word].update(comment=comment)

    def to_list(self) -> List[dict[str, Any]]:
        """
        Returns a list of objects. Each object contains a "filename" and "tags0" that are a list of words.
        """
        result = []
        for filename, tags in self.data.items():
            entry = {
                "filename": filename.tag,
                "title": filename.title,
                "tags0": list(sorted(tags.keys(), key=lambda x: len(x))),
            }
            tmp = set(
                part for tag in tags.keys() for part in tag.split("_") if "_" in tag
            ).union(tag.split(".")[-1] for tag in tags.keys() if "." in tag)
            tmp = set(
                part for part in tmp if len(part) >= 2 and not re.match(r"^\d+$", part)
            )
            entry["tags1"] = list(sorted(tmp))
            entry["tags2"] = list(
                word
                for k in tags.values()
                if k.comment
                for word in re.split(r"[\W_]+", k.comment)
                if len(word) >= 3 and not re.match(r"^\d+$", word)
            )
            result.append(entry)
        # Cleanup common words
        for entry in result:
            entry["tags2"] = " ".join((w for w in entry["tags2"])).strip()
        # Filter out common words
        return result


def _parseHints(description: str | None) -> Tuple[str, Dict[str, str]]:
    """
    Parse the hints in the description and return a list of tuples with the hint type and value.
    """
    hints = {}
    if not description:
        return "", hints
    result = []
    for line in description.splitlines():
        line = line.rstrip()
        # Detect lines starting with a single space and remove the space
        if RE_SINGLE_SPACE.match(line):
            line = line[1:]
        if (m := RE_HINT.match(line)) is not None:
            hints[m.group(1)] = m.group(2)
            line = RE_HINT.sub("", line)
            if line:
                result.append(line)
        else:
            result.append(line)

    return "\n".join(result).strip(), hints


def sanitizeUrl(url: str) -> str:
    return (
        url.replace(" ", "-")
        .replace("/", "-")
        .replace(".", "-")
        .replace("(", "")
        .replace(")", "")
        .lower()
    )


def getLinkFromType(
    full_type: str, clean_google_empty=False, subtype=None, enumList: List[str] = None
) -> Tuple[str, str]:
    if (m := re.match(r"^map<([^>,]+), ([^>]+)>$", full_type)) is not None:
        _, k = getLinkFromType(m.group(1), enumList=enumList)
        _, v = getLinkFromType(m.group(2), enumList=enumList)
        return full_type, f"map<{k}, {v}>"

    fn_prefix = "model-"
    if enumList and (full_type in enumList):
        fn_prefix = "enum-"

    full_type_link = ""
    if full_type == "google.protobuf.Empty" and clean_google_empty:
        full_type = ""
    elif full_type.startswith("google.") or "." not in full_type:
        # simple proto types does not contian a dot, so they are not linked
        # google.* types are not linked
        if subtype is not None:
            full_type_link = f"`{full_type} - {subtype}`"
        else:
            full_type_link = f"`{full_type}`"
    else:
        if subtype is not None:
            full_type_link = (
                f"[`{full_type} - {subtype}`]({fn_prefix}{sanitizeUrl(full_type)}.md)"
            )
        else:
            full_type_link = f"[`{full_type}`]({fn_prefix}{sanitizeUrl(full_type)}.md)"
    return full_type, full_type_link


def sanitize_search_words(tags: set) -> List[str]:
    tmp = []
    for t in tags:
        if t.startswith("io.clbs."):
            tmp.append(t.split(".")[-1])
        elif t.startswith("google.protobuf."):
            continue
        else:
            tmp.extend((i for i in re.split(r"\W+", t) if i))
    tags = set(t.lower() for t in tmp)
    return list(sorted(tags))


def generate(
    sable_config: SableConfig,
    sable_context: SableContext,
    visible_packages: List[Package],
    tagged_services: List[Tuple[Service, List[Tuple[str, List[Any]]]]],
):

    # FIXME:
    output_dir = os.path.join("gen", "markdown")

    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    search_index: SearchIndex = SearchIndex()

    with open(os.path.join(output_dir, "index.md"), "w") as fh:
        fh.write("# API\n\n")
        # Write a paragraph about the API services and models
        fh.write(
            "This page lists all the services and models available in the API. Each service has a list of methods, and each method has a request and response model. The models are described in detail on their respective pages.\n\n"
        )

        # Write a list of services
        for service, groups in tagged_services:
            fh.write(f"# {service.name}\n\n")
            if service.description:
                fh.write(f"{service.description}\n\n")
            for group, methods in groups:
                fh.write(
                    f"- [{group}](service-{sanitizeUrl(group)}-{sanitizeUrl(service.name)}.md)\n"
                )

    enum_list = list(set(e.full_name for p in visible_packages for e in p.enums))

    for service, groups in tagged_services:
        for group, methods in groups:
            filename = f"service-{sanitizeUrl(group)}-{sanitizeUrl(service.name)}.md"
            with open(
                os.path.join(
                    output_dir,
                    filename,
                ),
                "w",
            ) as fh:
                title = f"{service.name} - {group}"
                fh.write(f"# {title}\n\n")
                for method, tags in methods:
                    search_index.add(
                        filename,
                        [
                            method.name,
                            method.request.full_type,
                            method.response.full_type,
                        ],
                        title=title,
                    )
                    # Request model name (if not google.protobuf.Empty)
                    m_request, m_request_link = getLinkFromType(
                        method.request.full_type,
                        clean_google_empty=True,
                        enumList=enum_list,
                    )
                    # Response model name (if not google.protobuf.Empty)
                    m_response, m_response_link = getLinkFromType(
                        method.response.full_type,
                        clean_google_empty=True,
                        enumList=enum_list,
                    )

                    fh.write(f"## {method.name}\n\n")
                    method_description, hints = _parseHints(method.description)
                    if method_description:
                        fh.write(f"{method_description}\n\n")
                    d_request_in = hints.get("param", "")
                    if d_request_in:
                        fh.write(f"### Input\n{d_request_in}\n\n")
                    d_request_out = hints.get("return", "")
                    if d_request_out:
                        fh.write(f"### Output\n{d_request_out}\n\n")
                    if m_response:
                        fh.write(
                            f"```proto\n{method.name}({m_request}) returns ({m_response})\n```\n\n"
                        )
                        if m_request_link:
                            fh.write(f"- Input: {m_request_link}\n")
                        fh.write(f"- Output: {m_response_link}\n\n")
                    else:
                        fh.write(f"```proto\n{method.name}({m_request})\n```\n\n")
                        if m_request_link:
                            fh.write(f"- Input: {m_request_link}\n\n")

    # Generate files describing models
    for package in visible_packages:
        for e in package.enums:
            filename = f"enum-{sanitizeUrl(e.full_name)}.md"
            with open(
                os.path.join(output_dir, filename),
                "w",
            ) as fh:
                title = f"Enum: {e.full_name}"
                fh.write(f"# {title}\n\n")
                enum_description, section_hints = _parseHints(e.description)
                if enum_description:
                    fh.write(f"{enum_description}\n\n")
                search_index.add(filename, title, enum_description, title=title)
                fh.write("## Options\n\n")
                if enum_values := e.values:
                    if "sort" in section_hints:
                        enum_values = sorted(e.values, key=lambda x: x.name)

                    fh.write("| Value | Description |\n| --- | --- |\n")
                    for v in enum_values:
                        # Split the field.description by \n and join them with <br> to create new lines in the markdown table cell
                        tmp, hints = _parseHints(v.description)
                        value_description = "<br>".join(tmp.split("\n"))
                        search_index.add(filename, v.name, tmp)

                        if (tmp := hints.get("values", None)) is not None:
                            value_description = (
                                f"{value_description}<br><b>Options:</b> {tmp}"
                            )
                        if (tmp := hints.get("default", None)) is not None:
                            value_description = (
                                f"{value_description}<br><b>Default value:</b> {tmp}"
                            )
                        if (tmp := hints.get("example", None)) is not None:
                            value_description = (
                                f"{value_description}<br><b>Example:</b> {tmp}"
                            )

                        fh.write(f"| {v.name} | {value_description} |\n")
        for message in package.messages:
            filename = f"model-{sanitizeUrl(message.full_name)}.md"
            with open(
                os.path.join(output_dir, filename),
                "w",
            ) as fh:
                title = f"Model: {message.full_name}"
                fh.write(f"# {title}\n\n")
                if message.description:
                    fh.write(f"{message.description}\n\n")
                search_index.add(
                    filename, message.full_name, message.description, title=title
                )
                if message.fields:
                    fh.write("## Fields\n\n")
                    # Generate message.fields as a table; look for longest values and align them to genarate a markdown table
                    # Keep in mind that the field.description may contain multiple lines separated by \n so handle them to keep the generated markdown table valid!
                    # ... it may need to prefix all lines and merge them fit into single cell - The key is to use HTML line breaks (<br>) within the cell content to create new lines.
                    fh.write("| Field | Information |\n| --- | --- |\n")
                    for field in message.fields:
                        # Split the field.description by \n and join them with <br> to create new lines in the markdown table cell
                        tmp, hints = _parseHints(field.description)
                        field_description = "<br>".join(tmp.split("\n"))
                        search_index.add(filename, field.name, tmp)

                        if (tmp := hints.get("values", None)) is not None:
                            field_description = (
                                f"{field_description}<br><b>Values:</b> {tmp}"
                            )
                        if (tmp := hints.get("default", None)) is not None:
                            field_description = (
                                f"{field_description}<br><b>Default value:</b> {tmp}"
                            )
                        if (tmp := hints.get("example", None)) is not None:
                            field_description = (
                                f"{field_description}<br><b>Example:</b> {tmp}"
                            )

                        _, full_type_link = getLinkFromType(
                            field.full_type,
                            clean_google_empty=False,
                            subtype=hints.get("gqltype", None),
                            enumList=enum_list,
                        )
                        field_description = f"<b>Type:</b> {full_type_link}<br><b>Description:</b><br>{field_description}"

                        fh.write(f"| {field.name} | {field_description} |\n")
                    fh.write("\n")

    with open(os.path.join(output_dir, "search_index.json"), "w") as fh:
        json.dump(search_index.to_list(), fh, indent=2)
