#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import os, re
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Any

re_format = re.compile(r"map<string,\s*([^\s]+)\s*>")
re_hint = re.compile(r"@([a-z]+):\s*(.*)")

# Map of protobuf types to GraphQL types
type_map = {
    "string": "String",
    "int32": "Int",
    "float": "Float",
    "bool": "Boolean",
    #
    "int64": "Int64",
    "sint64": "Int64",
    "double": "Float",
    "bytes": "String",
    "uint32": "Int",
    "uint64": "BigInt",
}


def format_comment(s: str, n: int = 0, block: bool = False) -> str:
    if not s or not s.strip():
        return ""
    lines = s.splitlines()
    cnt = len(lines)
    if cnt == 1 and not block and not '"' in s:
        return " " * n + "# {}\n".format(s.strip())
    s = "\n".join((" " * n + "# " + line).rstrip() for line in lines)
    return s + "\n"


def parse_comment(s: str) -> str:
    m = re_hint.findall(s)
    if m:
        hints = {}
        for k, v in m:
            hints[k] = v
        return re_hint.sub("", s).strip(), hints
    return s.rstrip(), {}


def generate(
    sable_config: SableConfig,
    sable_context: SableContext,
    visible_packages: List[Package],
    tagged_services: List[Tuple[Service, List[Tuple[str, List[Any]]]]],
):

    # FIXME:
    output_dir = "graph"

    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    map_names = {}
    scalars = set()

    with open(os.path.join(output_dir, "api.graphqls"), "w") as fh:
        fh.write("# Code generated, DO NOT EDIT.\n")
        fh.write("\n")

        input_types = set()
        for _, groups in tagged_services:
            for _, methods in groups:
                for method, _ in methods:
                    if method.request.full_type in (
                        "google.protobuf.Empty",
                        "Timestamp",
                    ):
                        continue

                    # FIXME: Input types
                    # input_types.add(method.request.full_type)

        # Generate enums
        for package in sable_context.packages:
            for enum in package.enums:
                fh.write(format_comment(enum.description, block=True))
                fh.write("enum {} {{\n".format(enum.name))
                for value in enum.values:
                    if value.description:
                        fh.write(format_comment(value.description, 2))
                    fh.write("  {}\n".format(value.name))
                fh.write("}\n")
                fh.write("\n")

        # Generate types
        known_type_map = {}
        known_type_names = set()
        for package in sable_context.packages:
            for message in package.messages:
                if message.name == "Timestamp":
                    scalars.add("Timestamp")
                    continue

                proposed_type_name = message.name
                if proposed_type_name in known_type_names:
                    proposed_type_suffix = 1
                    while True:
                        tmp = "{}_{}".format(message.name, proposed_type_suffix)
                        if tmp not in known_type_names:
                            proposed_type_name = tmp
                            break
                        proposed_type_suffix += 1
                known_type_map[message.full_type] = proposed_type_name
                known_type_names.add(proposed_type_name)

                if message.full_type in input_types:
                    _generateMessage(
                        "input", message, proposed_type_name + "Input", map_names, fh
                    )
                _generateMessage("type", message, proposed_type_name, map_names, fh)

        # Proto map types need additional types in GraphQL
        # map<string, any> -> type _mapAny { key: String!, value: String }
        for map_type_name, map_type_value in map_names.items():
            if (mapped_type := type_map.get(map_type_value, None)) is not None:
                map_type_value = mapped_type
            fh.write("type {} {{\n".format(map_type_name))
            fh.write("  key: String!\n")
            fh.write("  value: {}\n".format(map_type_value))
            fh.write("}\n\n")

        # Generate queries
        fh.write("type Query {\n")
        for service, groups in tagged_services:
            for group, methods in groups:
                for method, tags in methods:
                    # camelCase method name
                    method_name = method.name[0].lower() + method.name[1:]
                    fh.write("  {}".format(method_name))
                    if method.request.full_type in input_types:
                        gql_input_type = (
                            known_type_map[method.request.full_type] + "Input"
                        )
                        fh.write("(data: {}): ".format(gql_input_type))
                    else:
                        fh.write(": ")
                    gql_output_type = known_type_map[method.response.full_type]
                    fh.write("{}\n".format(gql_output_type))
        fh.write("}\n\n")

        # Generate mutations
        # fh.write("type Mutation {\n")
        # for service, methods in tagged_services:
        #     for group, methods in groups:
        #         for method, _ in methods:
        #             method_name = method.name[0].lower() + method.name[1:]
        #             fh.write("  {}".format(method_name))
        #             if method.request.full_type:
        #                 fh.write(
        #                     "({}: {})".format(
        #                         method.request.name, method.request.full_type
        #                     )
        #                 )
        #             fh.write(": {}\n".format(method.response.full_type))
        # fh.write("}\n\n")

        # Generate scaler types
        for scalar in scalars:
            fh.write("scalar {}\n".format(scalar))
        fh.write("scalar BigInt\n")
        fh.write("scalar Int64\n")
        fh.write("scalar UUID\n")


def _generateMessage(object_def_name, message, type_name, map_names, fh):
    has_oneof = any(f.oneof_name for f in message.fields)

    description, hints = parse_comment(message.description)
    fh.write(format_comment(description, block=True))
    fh.write("{} {} {{\n".format(object_def_name, type_name))
    for field in message.fields:
        if field.oneof_name:
            # TODO: handle oneof
            pass
        description, hints = parse_comment(field.description)
        ftype = field.type
        if (mapped_type := hints.get("gqltype", None)) is not None:
            ftype = mapped_type
        elif (mapped_type := type_map.get(ftype, None)) is not None:
            ftype = mapped_type
        if field.label == "repeated":
            ftype = "[{}]".format(ftype)
        fh.write(format_comment(description, 2))
        if (m := re_format.match(ftype)) is not None:
            map_type_name = "_map{}".format(m.group(1))
            map_names[map_type_name] = m.group(1)
            ftype = "[{}]".format(map_type_name)
        fh.write("  {}: {}\n".format(field.name, ftype))
    if not message.fields:
        fh.write("  _empty: Boolean\n")
    fh.write("}\n\n")
