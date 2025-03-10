# Model: io.clbs.openhes.models.common.FieldDescriptor

The field descriptor.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| fieldId | string | Unique identifier for the field |
| label | string | Label displayed for the field |
| dataType | [io.clbs.openhes.models.common.FieldDataType](model-io-clbs-openhes-models-common-fielddatatype.md) | Data type of the field (e.g., text, double) |
| format | [io.clbs.openhes.models.common.FieldDisplayFormat](model-io-clbs-openhes-models-common-fielddisplayformat.md) | Display format (e.g., 1h 30m) |
| unit | string | Unit to display (e.g., kWh, USD) |
| groupId | string | Group (section) identifier for the field |
| precision | int32 | Decimal precision for double numbers |
| tooltip | string | Tooltip or hint text |
| required | bool | Whether the field is mandatory |
| editable | bool | Whether the field is editable |
| visible | bool | Whether the field is visible |
| multiValue | bool | Whether the field can have multiple values |
| secured | bool | Whether the field shall be handled as a security fields (e.g., password, certificate input area, ...) |
| validation | [io.clbs.openhes.models.common.FieldValidation](model-io-clbs-openhes-models-common-fieldvalidation.md) | Validation rules for the field |
| defaultValue | [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md) | The default value of the attribute, it does not support multi-value fields |
| jsPath | string | The path to the field in the JSON data model |

