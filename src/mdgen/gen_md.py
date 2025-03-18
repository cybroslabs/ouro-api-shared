#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, re
from shutil import copy, copytree, rmtree
from sabledocs.lunr_search import build_search_index
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Any


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


def getLinkFromType(full_type: str, clean_google_empty=False) -> Tuple[str, str]:
    if (m := re.match(r"^map<([^>,]+), ([^>]+)>$", full_type)) is not None:
        _, k = getLinkFromType(m.group(1))
        _, v = getLinkFromType(m.group(2))
        return full_type, f"map<{k}, {v}>"

    full_type_link = ""
    if full_type == "google.protobuf.Empty" and clean_google_empty:
        full_type = ""
    elif full_type.startswith("google.") or "." not in full_type:
        # simple proto types does not contian a dot, so they are not linked
        # google.* types are not linked
        full_type_link = full_type
    else:
        full_type_link = f"[{full_type}](model-{sanitizeUrl(full_type)}.md)"
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
                        method.request.full_type, clean_google_empty=True
                    )
                    # Response model name (if not google.protobuf.Empty)
                    m_response, m_response_link = getLinkFromType(
                        method.response.full_type, clean_google_empty=True
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
                            field_description = "<br>".join(
                                field.description.split("\n")
                            )
                        else:
                            field_description = ""
                        _, full_type_link = getLinkFromType(
                            field.full_type, clean_google_empty=False
                        )
                        field_description = f"<b>Type:</b> {full_type_link}<br><b>Description:</b><br>{field_description}"
                        fh.write(f"| {field.name} | {field_description} |\n")
                    fh.write("\n")
