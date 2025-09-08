# Model: io.clbs.openhes.models.common.MetadataFields

Defines a specification for metadata fields managed by both users and the system.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The UUID of the resource. Serves as the unique identifier of the resource. Immutable and typically auto-generated during `Create` operations. |
| generation | <b>Type:</b> `int32`<br><b>Description:</b><br>The resource generation. |
| fields | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>Additional fields managed by the user. |
| managedFields | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>Additional fields managed by the system. |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The resource name. Mutable and typically set by the user. Must be set. |

