# Model: io.clbs.openhes.models.common.FieldValidation

Validation rules for the field.

## Fields

| Field | Information |
| --- | --- |
| re | <b>Type:</b> `string`<br><b>Description:</b><br>Regular expression describing input format. If not set then any value of given type can be used. It can be used for string, int or double fields only. |
| minLength | <b>Type:</b> `int32`<br><b>Description:</b><br>The minimum length. It's used for string fields only. |
| maxLength | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum length. It's used for string fields only. |
| minInteger | <b>Type:</b> `sint64`<br><b>Description:</b><br>The minimum value. It's used for integer fields only. |
| maxInteger | <b>Type:</b> `sint64`<br><b>Description:</b><br>The maximum value. It's used for integer fields only. |
| minNumber | <b>Type:</b> `double`<br><b>Description:</b><br>The minimum value. It's used for number fields only. |
| maxNumber | <b>Type:</b> `double`<br><b>Description:</b><br>The maximum value. It's used for number fields only. |
| options | <b>Type:</b> `map<string, string>`<br><b>Description:</b><br>The list of allowed values to be set (key-value pairs). The key here represents the field value to be set and the value here represents the label to be displayed. |

