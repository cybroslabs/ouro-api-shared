#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, json, itertools, re, copy
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import Set, Dict, List, Any, Tuple
from proto_descriptor_parser import markdown_to_html

re_hint = re.compile(r"@([a-z]+): (.*)")
re_spaces = re.compile(r"\s+")
re_type_map = re.compile(r"map<string,\s*([^\s]+)\s*>")


def filter_used(
    packages_map, base_messages: Dict[str, str], base_enums: Dict[str, str]
):
    while True:
        # Clone before iterating
        done = True
        for t_name, p_name in list(base_messages.items()):
            if p_name not in packages_map:
                continue
            for p in packages_map[p_name].messages:
                if p.full_type == t_name:
                    if t_name.endswith("JobStatus"):
                        pass
                    for f in p.fields:
                        if f.type_kind == "ENUM":
                            if f.full_type not in base_enums:
                                # New enum
                                done = False
                                base_enums[f.full_type] = f.package.name

                        elif f.type_kind == "MESSAGE":
                            if f.package:
                                if f.full_type not in base_messages:
                                    # New type
                                    done = False
                                    base_messages[f.full_type] = f.package.name
                            elif (
                                ft := re_type_map.match(f.full_type or "")
                            ) is not None:
                                # This does not work as the map type is not containing the package name...
                                parts = ft.group(1).rsplit(".", 1)
                                if len(parts) == 2:
                                    if ft.group(1) not in base_messages:
                                        # New type
                                        done = False
                                        base_messages[ft.group(1)] = parts[0]
        if done:
            break


def run(
    sable_config: SableConfig, sable_context: SableContext
) -> Tuple[Set[Package], List[Tuple[Service, List[Tuple[str, List[Any]]]]]]:
    visible_packages = set(
        (p for p in sable_context.non_hidden_packages if p.name.startswith("io.clbs"))
    )

    package_map: Dict[str, Package] = {p.name: p for p in sable_context.packages}
    main_package = package_map["io.clbs.openhes.pbapi"]

    # Go though non-hidden packages and gather all the used messages and enums
    used_messages = {
        m.full_type: (m.package.name if m.package else None)
        for m in main_package.messages
    }
    used_enums = {
        e.full_name: (e.package.name if e.package else None) for e in main_package.enums
    }

    used_messages.update(
        (
            (
                n.full_type,
                n.package.name if n.package else None,
            )
            for m in main_package.messages
            for n in m.fields
        )
    )
    for svc in main_package.services:
        for m in svc.methods:
            if m.request.type_kind == "MESSAGE":
                used_messages[m.request.full_type] = (
                    m.request.package.name if m.request.package else None
                )
            elif m.request.type_kind == "ENUM":
                used_enums[m.request.full_name] = (
                    m.request.package.name if m.request.package else None
                )
            if m.response.type_kind == "MESSAGE":
                used_messages[m.response.full_type] = (
                    m.response.package.name if m.response.package else None
                )
            elif m.response.type_kind == "ENUM":
                used_enums[m.response.full_name] = (
                    m.response.package.name if m.response.package else None
                )

    filter_used(package_map, used_messages, used_enums)

    used_messages = set((k, v) for k, v in used_messages.items())
    used_enums = set((k, v) for k, v in used_enums.items())

    # Filter out only what is used
    for package in sable_context.packages:
        if package.name == main_package.name:
            continue

        package.services = []
        package.messages = [
            m
            for m in package.messages
            if any((m.full_type == full_type for full_type, _ in used_messages))
        ]
        package.enums = [
            e
            for e in package.enums
            if any((e.full_name == full_type for full_type, _ in used_enums))
        ]

        if not (package.services or package.messages or package.enums):
            visible_packages.discard(package)

    tagged_services = []
    for svc in sorted(main_package.services, key=lambda x: x.name):
        tagger_methods = []
        for i in svc.methods:
            m = re_hint.findall(i.description)
            if m:
                group = next((x[1] for x in m if x[0] == "group"), None)
                tags = list((re_spaces.sub("", x[1]) for x in m if x[0] == "tag"))
                hints = list((x for x in m))

                desc = re_hint.sub("", i.description).strip()
                i.description_html = markdown_to_html(desc, sable_config)

            else:
                group = None
                tags = []
                hints = []

            tagger_methods.append(
                {"group": group, "method": i, "tags": tags, "hints": hints}
            )
        tagged_groups = list(
            sorted(
                (
                    (a, list(((m["method"], m["tags"]) for m in b)))
                    for a, b in itertools.groupby(
                        tagger_methods, key=lambda x: x["group"]
                    )
                ),
                key=lambda x: x[0] if x[0] else "",
            )
        )
        tagged_services.append((svc, tagged_groups))

    return visible_packages, tagged_services
