#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, re
from shutil import copy, copytree, rmtree
from sabledocs.lunr_search import build_search_index
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Dict, Any

re_hint = re.compile(r"@([a-z]+): (.*)")


def _parseHints(description: str) -> Tuple[str, Dict[str, str]]:
    """
    Parse the hints in the description and return a list of tuples with the hint type and value.
    """
    result = []
    hints = {}
    for line in description.splitlines():
        if (m := re_hint.match(line)) is not None:
            hints[m.group(1)] = m.group(2)
            line = re_hint.sub("", line).strip()
            if line:
                result.append(line)
        else:
            result.append(line)

    return "\n".join(result).strip(), hints


def _sanitizeMethodDescripton(description: str) -> str:
    description = [
        line for line in description.split("\n") if not line.lstrip().startswith("@")
    ]
    return "\n".join(description).strip()


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
            with open(
                os.path.join(
                    output_dir,
                    f"service-{sanitizeUrl(group)}-{sanitizeUrl(service.name)}.md",
                ),
                "w",
            ) as fh:
                fh.write(f"# {service.name} - {group}\n\n")
                for method, tags in methods:
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
                    if (d := _sanitizeMethodDescripton(method.description)) and d:
                        fh.write(f"{d}\n\n")
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
            # model-io-clbs-openhes-models-common-fielddisplayformat.md
            with open(
                os.path.join(output_dir, f"enum-{sanitizeUrl(e.full_name)}.md"),
                "w",
            ) as fh:
                fh.write(f"# Enum: {e.full_name}\n\n")
                if e.description:
                    fh.write(f"{e.description}\n\n")
                fh.write("## Options\n\n")
                if e.values:
                    fh.write("| Value | Description |\n| --- | --- |\n")
                    for v in e.values:
                        if v.description:
                            # Split the field.description by \n and join them with <br> to create new lines in the markdown table cell
                            tmp, hints = _parseHints(v.description)
                            value_description = "<br>".join(tmp.split("\n"))
                        else:
                            hints = {}
                            value_description = ""

                        value_description = (
                            value_description.replace("@values:", "<b>Values:</b>")
                            .replace("@default:", "<b>Default value:</b>")
                            .replace("@example:", "<b>Example:</b>")
                        )
                        fh.write(f"| {v.name} | {value_description} |\n")
        for message in package.messages:
            with open(
                os.path.join(output_dir, f"model-{sanitizeUrl(message.full_name)}.md"),
                "w",
            ) as fh:
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
                        if field.description:
                            # Split the field.description by \n and join them with <br> to create new lines in the markdown table cell
                            tmp, hints = _parseHints(field.description)
                            field_description = "<br>".join(tmp.split("\n"))
                        else:
                            hints = {}
                            field_description = ""

                        field_description = (
                            field_description.replace("@values:", "<b>Values:</b>")
                            .replace("@default:", "<b>Default value:</b>")
                            .replace("@example:", "<b>Example:</b>")
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
