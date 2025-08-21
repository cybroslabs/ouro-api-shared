import sys
from sabledocs.proto_model import SableConfig, SableContext, Package, Service
from typing import List, Tuple, Any

# List of reserved words in Python
RESERVED_WORDS = (
    "False",
    "await",
    "else",
    "import",
    "pass",
    "None",
    "break",
    "except",
    "in",
    "raise",
    "True",
    "class",
    "finally",
    "is",
    "return",
    "and",
    "continue",
    "for",
    "lambda",
    "try",
    "as",
    "def",
    "from",
    "nonlocal",
    "while",
    "assert",
    "del",
    "global",
    "not",
    "with",
    "async",
    "elif",
    "if",
    "or",
    "yield",
)


def validate(
    visible_packages: List[Package],
    tagged_services: List[Tuple[Service, List[Tuple[str, List[Any]]]]],
) -> bool:
    lower_reserved_words = {word.lower() for word in RESERVED_WORDS}

    # Check if a message, enum or a fields name is a reserved word
    valid = True
    for package in visible_packages:
        for message in package.messages:
            if message.name in RESERVED_WORDS:
                print(
                    f"Error: Message name '{message.name}' is a reserved word.",
                    file=sys.stderr,
                )
                valid = False
            elif message.name.lower() in lower_reserved_words:
                print(
                    f"Warning: Message name '{message.name}' is a reserved word (case-insensitive).",
                    file=sys.stderr,
                )
            for field in message.fields:
                if field.name in RESERVED_WORDS:
                    print(
                        f"Error: Field name '{field.name}' in message '{message.name}' is a reserved word.",
                        file=sys.stderr,
                    )
                    valid = False
                elif field.name.lower() in lower_reserved_words:
                    print(
                        f"Warning: Field name '{field.name}' in message '{message.name}' is a reserved word (case-insensitive).",
                        file=sys.stderr,
                    )

        for enum in package.enums:
            if enum.name in RESERVED_WORDS:
                print(
                    f"Error: Enum name '{enum.name}' is a reserved word.",
                    file=sys.stderr,
                )
                valid = False
            elif enum.name.lower() in lower_reserved_words:
                print(
                    f"Warning: Enum name '{enum.name}' is a reserved word (case-insensitive).",
                    file=sys.stderr,
                )
            for value in enum.values:
                if value.name in RESERVED_WORDS:
                    print(
                        f"Error: Enum value '{value.name}' in enum '{enum.name}' is a reserved word.",
                        file=sys.stderr,
                    )
                    valid = False
                elif value.name.lower() in lower_reserved_words:
                    print(
                        f"Warning: Enum value '{value.name}' in enum '{enum.name}' is a reserved word (case-insensitive).",
                        file=sys.stderr,
                    )

    for service in tagged_services:
        if service[0].name in RESERVED_WORDS:
            print(
                f"Error: Service name '{service[0].name}' is a reserved word.",
                file=sys.stderr,
            )
            valid = False
        elif service[0].name.lower() in lower_reserved_words:
            print(
                f"Warning: Service name '{service[0].name}' is a reserved word (case-insensitive).",
                file=sys.stderr,
            )
        for method in service[1]:
            if method[0] in RESERVED_WORDS:
                print(
                    f"Error: Method name '{method[0]}' in service '{service[0].name}' is a reserved word.",
                    file=sys.stderr,
                )
                valid = False
            elif method[0].lower() in lower_reserved_words:
                print(
                    f"Warning: Method name '{method[0]}' in service '{service[0].name}' is a reserved word (case-insensitive).",
                    file=sys.stderr,
                )

    return valid
