# Model: io.clbs.openhes.models.acquisition.MeasuredTimestampTzValue

Defines a measured value.

## Fields

| Field | Information |
| --- | --- |
| status | <b>Type:</b> `int64`<br><b>Description:</b><br>The status of the value. |
| exponent | <b>Type:</b> `int32`<br><b>Description:</b><br>The exponent of the value. Used to calculate the real value as `value * 10^exponent` for double and integer values. |
| timestampTzValue | <b>Type:</b> `string`<br><b>Description:</b><br>The timestamp with timezone-typed value, stored as an ISO-8601 string. |
| nstatus | <b>Type:</b> `uint64`<br><b>Description:</b><br>The normalized status of the value. See `StatusBits` enum for details. |
| peakTs | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the value (for example, peak) was recorded. Typically before the capture timestamp. |

