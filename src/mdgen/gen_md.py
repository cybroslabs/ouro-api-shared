#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, re, json
from shutil import copy, copytree, rmtree
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Dict, Any

RE_HINT = re.compile(r"\s*@([a-z]+)(?::?\s*(.*))?$")
RE_SINGLE_SPACE = re.compile(r"^\s(?:[^\s]*|$)")


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

    search_index: dict[str, set] = {}

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
            search_index_page = search_index[filename] = set()
            with open(
                os.path.join(
                    output_dir,
                    filename,
                ),
                "w",
            ) as fh:
                search_index_page.update([service.name, group])
                fh.write(f"# {service.name} - {group}\n\n")
                for method, tags in methods:
                    search_index_page.update(
                        [
                            method.name,
                            method.request.full_type,
                            method.response.full_type,
                        ]
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
            search_index_page = search_index[filename] = set()
            with open(
                os.path.join(output_dir, filename),
                "w",
            ) as fh:
                search_index_page.add(e.full_name)
                fh.write(f"# Enum: {e.full_name}\n\n")
                enum_description, section_hints = _parseHints(e.description)
                if enum_description:
                    fh.write(f"{enum_description}\n\n")
                fh.write("## Options\n\n")
                if enum_values := e.values:
                    if "sort" in section_hints:
                        enum_values = sorted(e.values, key=lambda x: x.name)

                    fh.write("| Value | Description |\n| --- | --- |\n")
                    for v in enum_values:
                        search_index_page.add(v.name)
                        # Split the field.description by \n and join them with <br> to create new lines in the markdown table cell
                        tmp, hints = _parseHints(v.description)
                        value_description = "<br>".join(tmp.split("\n"))

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
            search_index_page = search_index[filename] = set()
            with open(
                os.path.join(output_dir, filename),
                "w",
            ) as fh:
                search_index_page.add(message.full_name)
                fh.write(f"# Model: {message.full_name}\n\n")
                if message.description:
                    fh.write(f"{message.description}\n\n")
                if message.fields:
                    fh.write("## Fields\n\n")
                    # Generate message.fields as a table; look for longest values and align them to genarate a markdown table
                    # Keep in mind that the field.description may contain multiple lines separated by \n so handle them to keep the generated markdown table valid!
                    # ... it may need to prefix all lines and merge them fit into single cell - The key is to use HTML line breaks (<br>) within the cell content to create new lines.
                    fh.write("| Field | Information |\n| --- | --- |\n")
                    for field in message.fields:
                        search_index_page.add(field.name)
                        # Split the field.description by \n and join them with <br> to create new lines in the markdown table cell
                        tmp, hints = _parseHints(field.description)
                        field_description = "<br>".join(tmp.split("\n"))

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
        index = [
            {"filename": k, "tags": sanitize_search_words(v)}
            for k, v in search_index.items()
        ]
        json.dump(index, fh, indent=2)
