# Model: io.clbs.openhes.models.common.FieldDescriptor

Defines the field descriptor specification.

## Fields

| Field | Information |
| --- | --- |
| isUserDefined | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field descriptor is user-defined (for example, a custom field added by users). |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>The resource type for which the field descriptor is defined (for example, `BULK`, `DEVICE`). |
| gid | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the field descriptor within the system. |
| fieldId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique field descriptor identifier. |
| jsPath | <b>Type:</b> `string`<br><b>Description:</b><br>The path to the field in the TypeScript representation. |
| path | <b>Type:</b> `string`<br><b>Description:</b><br>The path to the field in the gRPC/JSON simplified representation. |
| label | <b>Type:</b> `string`<br><b>Description:</b><br>The label displayed for the field. |
| groupId | <b>Type:</b> `string`<br><b>Description:</b><br>The group (section) identifier for the field. |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the field (for example, `text`, `double`). |
| format | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDisplayFormat`](enum-io-clbs-openhes-models-common-fielddisplayformat.md)<br><b>Description:</b><br>The display format of the field (for example, `1h 30m`). |
| unit | <b>Type:</b> `string`<br><b>Description:</b><br>The display unit for the field (fr example, `kWh`, `USD`). |
| precision | <b>Type:</b> `int32`<br><b>Description:</b><br>The decimal precision for double values. |
| tooltip | <b>Type:</b> `string`<br><b>Description:</b><br>A tooltip or hint text for tehe field. |
| required | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field is mandatory. |
| editable | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field is editable. |
| visible | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field is visible. |
| multiValue | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field can contain multiple values. |
| secured | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field should be handled as security fields (for example, password or certificate input area). |
| validation | <b>Type:</b> [`io.clbs.openhes.models.common.FieldValidation`](model-io-clbs-openhes-models-common-fieldvalidation.md)<br><b>Description:</b><br>The validation rules for the field. |
| defaultValue | <b>Type:</b> [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)<br><b>Description:</b><br>The default value of the . Multi-value fields are not supported. |
| sortable | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field can be used for sorting. Default is `true`. |
| filterable | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the field can be used for filtering. Default is `true`. |

