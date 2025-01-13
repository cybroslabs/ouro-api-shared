#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os
import json
import markdown
from shutil import copy, copytree, rmtree
from jinja2 import Environment, FileSystemLoader, select_autoescape
from sabledocs.lunr_search import build_search_index
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Any


def generate(
    sable_config: SableConfig,
    sable_context: SableContext,
    visible_packages: List[Package],
    tagged_services: List[Tuple[Service, List[Tuple[str, List[Any]]]]],
):
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
                tagger_services=tagged_services,
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
