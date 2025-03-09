# Model: io.clbs.openhes.models.acquisition.MeasuredValue

Sub-message containing measured value

## Fields

| Field | Type | Description |
| --- | --- | --- |
| status | int64 | The status of the value. |
| exponent | int32 | The exponent of the value. It's used to calculate the real value as value * 10^exponent for double and int values. |
| doubleValue | double | The double-typed value. |
| integerValue | int64 | The integer-typed value. |
| stringValue | string | The string-typed value. |
| timestampValue | google.protobuf.Timestamp | The timestamp-typed value. |
| timestampTzValue | string | The timestamp with timezone-typed value, stored as string in ISO-8601. |
| boolValue | bool | The boolean-typed value. |

