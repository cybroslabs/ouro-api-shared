# Model: io.clbs.openhes.models.common.SetFieldsSpec

Defines a specification for creating or updating fields for resources.

## Fields

| Field | Information |
| --- | --- |
| fields | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The fields to set for the resources. |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>The type of the managed fields update. Defines resource type the fields apply to. |
| objectId | <b>Type:</b> `string`<br><b>Description:</b><br>The UUID of the resource. Serves as the unique identifier of the resource. Immutable and typically auto-generated during `Create` operations. |

