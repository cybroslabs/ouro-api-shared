#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, json, itertools, re
import markdown
from shutil import copy, copytree, rmtree
from proto_descriptor_parser import (
    parse_proto_descriptor,
    SableConfig,
    markdown_to_html,
)
from sabledocs.lunr_search import build_search_index
from jinja2 import Environment, FileSystemLoader, select_autoescape
from typing import Set

re_group = re.compile(r"@([a-z]+): (.*)")
re_spaces = re.compile(r"\s+")


def filter_used(packages_map, base_messages: Set[str], base_enums: Set[str]):
    # Clone before iterating
    nested_packages = set()
    for t_name, p_name in list(base_messages):
        if p_name not in packages_map:
            continue
        for p in packages_map[p_name].messages:
            if p.full_type == t_name:
                for f in p.fields:
                    if f.package:
                        base_messages.add((f.full_type, f.package.name))
                        if f.package.name != p_name:
                            nested_packages.add(f.package.name)

    package_map = {k: v for k, v in packages_map.items() if k in nested_packages}
    if package_map:
        filter_used(package_map, base_messages, base_enums)


if __name__ == "__main__":
    sable_config = SableConfig("mdgen.toml")
    sable_config.input_descriptor_file = os.path.join("protobuf", "pbapi", "pbapi.pb")
    sable_context = parse_proto_descriptor(sable_config)

    template_base_dir = (
        sable_config.template_path
        if sable_config.template_path
        else os.path.join(os.path.dirname(__file__), "templates", "_default")
    )

    jinja_env = Environment(
        loader=FileSystemLoader(searchpath=template_base_dir),
        autoescape=select_autoescape(),
    )

    package_template = jinja_env.get_template("package.html")

    if not os.path.exists(sable_config.output_dir):
        os.makedirs(sable_config.output_dir)

    visible_packages = set(
        (p for p in sable_context.non_hidden_packages if p.name.startswith("io.clbs"))
    )

    package_map = {p.name: p for p in sable_context.packages}
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

    for package in visible_packages:
        with open(
            os.path.join(
                sable_config.output_dir,
                f'{package.name if package.name else "__default"}.html',
            ),
            "wb",
        ) as fh:
            output = package_template.render(
                sable_config=sable_config,
                package=package,
                packages=sable_context.packages,
                non_hidden_packages=visible_packages,
                all_messages=sable_context.all_messages,
                all_enums=sable_context.all_enums,
            ).encode("utf-8")

            fh.write(output)

    main_page_content = ""

    if sable_config.main_page_content_file != "":
        print()
        if os.path.exists(sable_config.main_page_content_file):
            print(f"Found main content page, {sable_config.main_page_content_file}.")
            with open(
                sable_config.main_page_content_file, mode="r"
            ) as main_page_content_file:
                main_page_content = markdown.markdown(
                    main_page_content_file.read(),
                    extensions=sable_config.markdown_extensions,
                )
        else:
            print(
                f"WARNING: The configured main content page, {sable_config.main_page_content_file} was not found."
            )

    tagger_services = []
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
        tagger_services.append((svc, tagged_groups))

    with open(os.path.join(sable_config.output_dir, "index.html"), "wb") as fh:
        output = (
            jinja_env.get_template("index.html")
            .render(
                sable_config=sable_config,
                main_page_content=main_page_content,
                packages=sable_context.packages,
                non_hidden_packages=visible_packages,
                all_messages=sable_context.all_messages,
                all_enums=sable_context.all_enums,
                tagger_services=tagger_services,
            )
            .encode("utf-8")
        )

        fh.write(output)

    if sable_config.enable_lunr_search:
        (search_documents, search_index) = build_search_index(sable_context)

        with open(os.path.join(sable_config.output_dir, "search.html"), "wb") as fh:
            output = (
                jinja_env.get_template("search.html")
                .render(
                    sable_config=sable_config,
                    search_documents=json.dumps(search_documents),
                    search_index=json.dumps(search_index.serialize()),
                )
                .encode("utf-8")
            )

            fh.write(output)

    index_abs_path = os.path.abspath(
        os.path.join(sable_config.output_dir, "index.html")
    )

    output_static_path = os.path.join(sable_config.output_dir, "static")

    if os.path.exists(output_static_path):
        # This is needed, because shutils.copytree cannot copy to a target folder which already exists.
        rmtree(output_static_path)

    copytree(os.path.join(template_base_dir, "static"), output_static_path)

    if os.path.exists("static") and os.path.isdir("static"):
        print("Copying static content from the folder 'static'.")
        for root, _, files in os.walk("static"):
            dir_path = "" if root == "static" else root.removeprefix("static\\")
            dest_dir_path = os.path.join(sable_config.output_dir, dir_path)
            if dir_path != "":
                if not os.path.exists(dest_dir_path):
                    os.makedirs(dest_dir_path)

            for f in files:
                src_file_path = os.path.join(root, f)
                dest_file_path = os.path.join(dest_dir_path, f)
                if not os.path.exists(dest_file_path):
                    copy(src_file_path, dest_file_path)

    if sable_config.extra_template_path != "":
        print(
            f"Rendering extra Jinja templates from, {sable_config.extra_template_path}"
        )
        jinja_extra_env = Environment(
            loader=FileSystemLoader(searchpath=sable_config.extra_template_path),
            autoescape=select_autoescape(),
        )
        for root, _, files in os.walk(sable_config.extra_template_path):
            dir_path = (
                ""
                if root == sable_config.extra_template_path
                else (root.removeprefix(sable_config.extra_template_path).rstrip("/\\"))
            )
            if (
                "/_" in dir_path or "\\_" in dir_path
            ):  # ignore subdirectories that start with "_"
                continue
            for file in files:
                if not file.endswith(sable_config.extra_template_suffix):
                    continue
                file_path = (
                    file if dir_path == "" else str(os.path.join(dir_path, file))
                )
                print(f"Rendering extra Jinja template, {file_path}")
                with open(os.path.join(sable_config.output_dir, file_path), "wb") as fh:
                    output = (
                        jinja_extra_env.get_template(file_path)
                        .render(sable_config=sable_config)
                        .encode("utf-8")
                    )

                    fh.write(output)
    print()
    print(f"Building documentation done. It can be opened with {index_abs_path}")
