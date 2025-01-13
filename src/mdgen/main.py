#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, json, itertools, re, copy
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

from processor import run
import gen_html

if __name__ == "__main__":
    sable_config = SableConfig("mdgen.toml")
    sable_config.input_descriptor_file = os.path.join(
        "protobuf", "pbapi", "pbapi.binpb"
    )
    sable_context = parse_proto_descriptor(sable_config)

    visible_packages, tagged_service = run(sable_config, sable_context)
    gen_html.generate(sable_config, sable_context, visible_packages, tagged_service)
