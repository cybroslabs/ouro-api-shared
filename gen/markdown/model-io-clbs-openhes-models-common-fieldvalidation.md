# Model: io.clbs.openhes.models.common.FieldValidation

Validation rules for the field.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| re | string | Regular expression describing input format. If not set then any value of given type can be used. It can be used for string, int or double fields only. |
| minLength | int32 | The minimum length. It's used for string fields only. |
| maxLength | int32 | The maximum length. It's used for string fields only. |
| minInteger | sint64 | The minimum value. It's used for integer fields only. |
| maxInteger | sint64 | The maximum value. It's used for integer fields only. |
| minNumber | double | The minimum value. It's used for number fields only. |
| maxNumber | double | The maximum value. It's used for number fields only. |
| options | map<string, string> | The list of allowed values to be set (key-value pairs). The key here represents the field value to be set and the value here represents the label to be displayed. |

