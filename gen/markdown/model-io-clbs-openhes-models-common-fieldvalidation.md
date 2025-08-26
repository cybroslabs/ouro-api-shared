# Model: io.clbs.openhes.models.common.FieldValidation

Defines validation rules applied to a field.

## Fields

| Field | Information |
| --- | --- |
| re | <b>Type:</b> `string`<br><b>Description:</b><br>A regular expression describing the allowed input format. If not set, any value of the given type can be used. Used for `string`, `integer` or `double` fields. |
| minLength | <b>Type:</b> `int32`<br><b>Description:</b><br>The minimum length of the field. Used for `string` fields. |
| maxLength | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum length of the field. Used for `string` fields. |
| minInteger | <b>Type:</b> `sint64`<br><b>Description:</b><br>The minimum value of the field. Used for `integer` fields.. |
| maxInteger | <b>Type:</b> `sint64`<br><b>Description:</b><br>The maximum value of the field. Used for `integer` fields. |
| minNumber | <b>Type:</b> `double`<br><b>Description:</b><br>The minimum value of the field. Used for `number` fields. |
| maxNumber | <b>Type:</b> `double`<br><b>Description:</b><br>The maximum value of the field. Used for `number` fields. |
| options | <b>Type:</b> map<`string`, `string`><br><b>Description:</b><br>A predefined list of allowed values to be set (as key-value pairs). The key represents the actual field value; the value represents the label to be displayed. |
| optionsSource | <b>Type:</b> `string`<br><b>Description:</b><br>The source from which the allowed options are dynamically fetched from the server, if set. |

