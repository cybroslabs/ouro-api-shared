#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os
from proto_descriptor_parser import (
    parse_proto_descriptor,
    SableConfig,
)

from processor import run
import gen_html
import gen_graphql
import gen_md

if __name__ == "__main__":
    sable_config = SableConfig("mdgen.toml")
    sable_config.input_descriptor_file = os.path.join(
        "gen", "go", "services", "svcapi", "raw.binpb"
    )
    sable_context = parse_proto_descriptor(sable_config)

    visible_packages, tagged_service = run(sable_config, sable_context)
    gen_html.generate(sable_config, sable_context, visible_packages, tagged_service)
    gen_graphql.generate(sable_config, sable_context, visible_packages, tagged_service)
    gen_md.generate(sable_config, sable_context, visible_packages, tagged_service)
