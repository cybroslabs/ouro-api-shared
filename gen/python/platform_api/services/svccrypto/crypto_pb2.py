# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svccrypto/crypto.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'services/svccrypto/crypto.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from crypto import crypto_pb2 as crypto_dot_crypto__pb2
from crypto import management_pb2 as crypto_dot_management__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fservices/svccrypto/crypto.proto\x12\"io.clbs.openhes.services.svccrypto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x13\x63rypto/crypto.proto\x1a\x17\x63rypto/management.proto2\xc4\x02\n\rCryproService\x12Y\n\x04\x44lms\x12%.io.clbs.openhes.models.crypto.DlmsIn\x1a&.io.clbs.openhes.models.crypto.DlmsOut(\x01\x30\x01\x12v\n\x0fGetCryptoSecret\x12\x35.io.clbs.openhes.models.crypto.GetCryptoSecretRequest\x1a,.io.clbs.openhes.models.crypto.CryptoSecrets\x12`\n\x0fSetCryptoSecret\x12\x35.io.clbs.openhes.models.crypto.SetCryptoSecretRequest\x1a\x16.google.protobuf.EmptyBAZ?github.com/cybroslabs/ouro-api-shared/gen/go/services/svccryptob\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svccrypto.crypto_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z?github.com/cybroslabs/ouro-api-shared/gen/go/services/svccrypto'
  _globals['_CRYPROSERVICE']._serialized_start=180
  _globals['_CRYPROSERVICE']._serialized_end=504
# @@protoc_insertion_point(module_scope)
