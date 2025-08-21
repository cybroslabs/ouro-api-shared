#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, argparse, sys
from proto_descriptor_parser import (
    parse_proto_descriptor,
    SableConfig,
)
from processor import run


def parse_args():
    parser = argparse.ArgumentParser(
        description="Generate documentation from protobuf descriptors."
    )
    parser.add_argument(
        "-m",
        "--mode",
        type=str,
        nargs="+",
        choices=["html", "graphql", "markdown"],
        help="Modes of operation: html, graphql, markdown.",
        required=True,
    )
    parser.add_argument(
        "--validate",
        action="store_true",
        help="Validate the protobuf descriptor file.",
    )
    return parser.parse_args()


if __name__ == "__main__":
    args = parse_args()
    modes = args.mode or []

    sable_config = SableConfig("mdgen.toml")
    sable_config.input_descriptor_file = os.path.join(
        "gen", "go", "services", "svcapi", "raw.binpb"
    )
    sable_context = parse_proto_descriptor(sable_config)

    visible_packages, tagged_service = run(sable_config, sable_context)

    if args.validate:
        import validator

        is_valid = validator.validate(visible_packages, tagged_service)
        if not is_valid:
            sys.exit(1)

    if "html" in modes:
        import gen_html

        gen_html.generate(sable_config, sable_context, visible_packages, tagged_service)

    if "graphql" in modes:
        import gen_graphql

        gen_graphql.generate(
            sable_config, sable_context, visible_packages, tagged_service
        )

    if "markdown" in modes:
        import gen_md

        gen_md.generate(sable_config, sable_context, visible_packages, tagged_service)
