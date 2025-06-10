# Model: io.clbs.openhes.models.common.FieldDescriptor

The field descriptor.

## Fields

| Field | Information |
| --- | --- |
| isUserDefined | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the field descriptor is user-defined (e.g., custom fields added by users) |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](model-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>Defines the resource type for which the field descriptor is defined, e.g., BULK, DEVICE, etc. |
| gid | <b>Type:</b> `string`<br><b>Description:</b><br>The system-wide unique identifier of the field descriptor. |
| fieldId | <b>Type:</b> `string`<br><b>Description:</b><br>Unique identifier for the field descriptor |
| jsPath | <b>Type:</b> `string`<br><b>Description:</b><br>The path to the field in the TypeScript representation |
| path | <b>Type:</b> `string`<br><b>Description:</b><br>The path to the field in the gRPC/JSON simplified representation |
| label | <b>Type:</b> `string`<br><b>Description:</b><br>Label displayed for the field |
| groupId | <b>Type:</b> `string`<br><b>Description:</b><br>Group (section) identifier for the field |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](model-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>Data type of the field (e.g., text, double) |
| format | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDisplayFormat`](model-io-clbs-openhes-models-common-fielddisplayformat.md)<br><b>Description:</b><br>Display format (e.g., 1h 30m) |
| unit | <b>Type:</b> `string`<br><b>Description:</b><br>Unit to display (e.g., kWh, USD) |
| precision | <b>Type:</b> `int32`<br><b>Description:</b><br>Decimal precision for double numbers |
| tooltip | <b>Type:</b> `string`<br><b>Description:</b><br>Tooltip or hint text |
| required | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the field is mandatory |
| editable | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the field is editable |
| visible | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the field is visible |
| multiValue | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the field can have multiple values |
| secured | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the field shall be handled as a security fields (e.g., password, certificate input area, ...) |
| validation | <b>Type:</b> [`io.clbs.openhes.models.common.FieldValidation`](model-io-clbs-openhes-models-common-fieldvalidation.md)<br><b>Description:</b><br>Validation rules for the field |
| defaultValue | <b>Type:</b> [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)<br><b>Description:</b><br>The default value of the attribute, it does not support multi-value fields |

