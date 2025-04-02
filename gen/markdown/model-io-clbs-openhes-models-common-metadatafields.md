# Model: io.clbs.openhes.models.common.MetadataFields

The metadata fields managed by user and system.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The UUID of the resource. It serves as the unique identifier of the resource. It's immutable and typically auto-generated during Create operations. |
| generation | <b>Type:</b> `int32`<br><b>Description:</b><br>The generation of the resource. |
| fields | <b>Type:</b> `map<string, [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md)>`<br><b>Description:</b><br>The additional fields managed by user. |
| managedFields | <b>Type:</b> `map<string, [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md)>`<br><b>Description:</b><br>The additional fields managed by system. |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The name of the resource. It's mutable and typically set by user. Must be set. |

