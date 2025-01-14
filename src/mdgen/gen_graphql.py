#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, re
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Any

re_format = re.compile(r"map<string,\s*([^\s]+)\s*>")

# Map of protobuf types to GraphQL types
typ_map = {
    "string": "String",
    "int32": "Int",
    "float": "Float",
    "bool": "Boolean",
}


def _padleft(s: str, n: int) -> str:
    return "\n".join((" " * n + line).rstrip() for line in s.splitlines())


def generate(
    sable_config: SableConfig,
    sable_context: SableContext,
    visible_packages: List[Package],
    tagged_services: List[Tuple[Service, List[Tuple[str, List[Any]]]]],
):
    if not os.path.exists(sable_config.output_dir):
        os.makedirs(sable_config.output_dir)

    map_names = {}

    with open(os.path.join(sable_config.output_dir, "api.graphqls"), "w") as fh:
        fh.write('"""\nCode generated, DO NOT EDIT.\n"""\n')
        fh.write("\n")
        for package in sable_context.packages:
            for enum in package.enums:
                if enum.description:
                    fh.write('"""\n{}\n"""\n'.format(enum.description)),
                fh.write("enum {} {{\n".format(enum.name))
                for value in enum.values:
                    if value.description:
                        fh.write(
                            '  """\n{}\n  """\n'.format(_padleft(value.description, 2))
                        ),
                    fh.write("  {}\n".format(value.name))
                fh.write("}\n")
                fh.write("\n")

        for package in sable_context.packages:
            for message in package.messages:
                if message.description:
                    fh.write('"""\n{}\n"""\n'.format(message.description)),
                fh.write("type {} {{\n".format(message.name))
                for field in message.fields:
                    ftype = field.type
                    if (mapped_type := typ_map.get(ftype, None)) is not None:
                        ftype = mapped_type
                    if field.label == "repeated":
                        ftype = "[{}]".format(ftype)
                    if field.description:
                        fh.write(
                            '  """\n{}\n  """\n'.format(_padleft(field.description, 2))
                        ),
                    if (m := re_format.match(ftype)) is not None:
                        map_type_name = "_Map{}".format(m.group(1))
                        map_names[map_type_name] = m.group(1)
                        ftype = "[{}]".format(map_type_name)
                    fh.write("  {}: {}\n".format(field.name, ftype))
                fh.write("}\n\n")

        for map_type_name, map_type_value in map_names.items():
            if (mapped_type := typ_map.get(map_type_value, None)) is not None:
                map_type_value = mapped_type
            fh.write("type {} {{\n".format(map_type_name))
            fh.write("  key: String!\n")
            fh.write("  value: {}\n".format(map_type_value))
            fh.write("}\n\n")
