# Model: io.clbs.openhes.models.common.FormattedMessage

Defines a structure of a user-facing formatted message.

## Fields

| Field | Information |
| --- | --- |
| message | <b>Type:</b> `string`<br><b>Description:</b><br>The message text or format string. |
| args | <b>Type:</b> `string`<br><b>Description:</b><br>A list of format arguments. If provided, the `message` field is used interpreted as a format string. |
| params | <b>Type:</b> `google.protobuf.Struct`<br><b>Description:</b><br>Additional named parameters as a single-level map. Nested objects or arrays are not allowed. |

