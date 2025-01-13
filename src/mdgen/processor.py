#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, json, itertools, re, copy
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import Set, Dict, List, Any, Tuple
from proto_descriptor_parser import markdown_to_html

re_group = re.compile(r"@([a-z]+): (.*)")
re_spaces = re.compile(r"\s+")
re_type_map = re.compile(r"map<string,\s*([^\s]+)\s*>")


def filter_used(packages_map, base_messages: Set[str], base_enums: Set[str]):
    origin = set(packages_map)
    # Clone before iterating
    nested_packages = set()
    for t_name, p_name in list(base_messages):
        if p_name not in packages_map:
            continue
        for p in packages_map[p_name].messages:
            if p.full_type == t_name:
                if t_name.endswith("GetDevicesCommunicationUnitsResponse"):
                    pass
                for f in p.fields:
                    if f.package:
                        base_messages.add((f.full_type, f.package.name))
                        if f.package.name != p_name:
                            nested_packages.add(f.package.name)
                    elif (ft := re_type_map.match(f.full_type or "")) is not None:
                        # This does not work as the map type is not containing the package name...
                        parts = ft.group(1).rsplit(".", 1)
                        if len(parts) == 2:
                            base_messages.add((ft.group(1), parts[0]))
                            if parts[0] != p_name:
                                nested_packages.add(parts[0])

    nested_packages_map = {
        k: v for k, v in packages_map.items() if k in nested_packages
    }
    if nested_packages_map:
        filter_used(nested_packages_map, base_messages, base_enums)


def run(
    sable_config: SableConfig, sable_context: SableContext
) -> Tuple[Set[Package], List[Tuple[Service, List[Tuple[str, List[Any]]]]]]:
    visible_packages = set(
        (p for p in sable_context.non_hidden_packages if p.name.startswith("io.clbs"))
    )

    package_map: Dict[str, Package] = {p.name: p for p in sable_context.packages}
    main_package = package_map["io.clbs.openhes.pbapi"]

    # Go though non-hidden packages and gather all the used messages and enums
    used_messages = set(
        (
            (m.full_type, m.package.name if m.package else None)
            for m in main_package.messages
        )
    )
    used_messages.update(
        (
            (
                n.full_type,
                n.package.name,
            )
            for m in main_package.messages
            for n in m.fields
            if n.package
        )
    )
    for svc in main_package.services:
        for m in svc.methods:
            used_messages.add(
                (
                    m.request.full_type,
                    m.request.package.name if m.request.package else None,
                )
            )
            used_messages.add(
                (
                    m.response.full_type,
                    m.response.package.name if m.response.package else None,
                )
            )

    used_enums = set(
        (
            (
                e.full_name,
                e.package.name if e.package else None,
            )
            for e in main_package.enums
        )
    )

    filter_used(package_map, used_messages, used_enums)

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
            m = re_group.findall(i.description)
            if m:
                group = next((x[1] for x in m if x[0] == "group"), None)
                tags = list((re_spaces.sub("", x[1]) for x in m if x[0] == "tag"))

                desc = re_group.sub("", i.description).strip()
                i.description_html = markdown_to_html(desc, sable_config)

            else:
                group = None

            tagger_methods.append({"group": group, "method": i, "tags": tags})
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
