# Model: io.clbs.openhes.models.common.MetadataFields

The metadata fields managed by user and system.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| id | string | The UUID of the resource. It serves as the unique identifier of the resource. It's immutable and typically auto-generated during Create operations. |
| generation | int32 | The generation of the resource. |
| fields | map<string, io.clbs.openhes.models.common.FieldValue> | The additional fields managed by user. |
| managedFields | map<string, io.clbs.openhes.models.common.FieldValue> | The additional fields managed by system. |
| name | string | The name of the resource. It's mutable and typically set by user. Must be set. |

