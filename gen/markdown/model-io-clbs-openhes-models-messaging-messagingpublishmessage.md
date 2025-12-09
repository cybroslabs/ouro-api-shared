# Model: io.clbs.openhes.models.messaging.MessagingPublishMessage

Defines a publish action for sending messages.

## Fields

| Field | Information |
| --- | --- |
| subject | <b>Type:</b> `string`<br><b>Description:</b><br>The subject of the message. The subject must start with the 'events.custom.' prefix.<br><b>Example:</b> "events.custom.sapmessage" |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The message data. |

