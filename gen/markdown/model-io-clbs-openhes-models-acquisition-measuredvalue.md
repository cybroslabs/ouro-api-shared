# Model: io.clbs.openhes.models.acquisition.MeasuredValue

Sub-message containing measured value

## Fields

| Field | Information |
| --- | --- |
| status | <b>Type:</b> int64<br><b>Description:</b><br>The status of the value. |
| exponent | <b>Type:</b> int32<br><b>Description:</b><br>The exponent of the value. It's used to calculate the real value as value * 10^exponent for double and int values. |
| doubleValue | <b>Type:</b> double<br><b>Description:</b><br>The double-typed value. |
| integerValue | <b>Type:</b> int64<br><b>Description:</b><br>The integer-typed value. |
| stringValue | <b>Type:</b> string<br><b>Description:</b><br>The string-typed value. |
| timestampValue | <b>Type:</b> google.protobuf.Timestamp<br><b>Description:</b><br>The timestamp-typed value. |
| timestampTzValue | <b>Type:</b> string<br><b>Description:</b><br>The timestamp with timezone-typed value, stored as string in ISO-8601. |
| boolValue | <b>Type:</b> bool<br><b>Description:</b><br>The boolean-typed value. |

