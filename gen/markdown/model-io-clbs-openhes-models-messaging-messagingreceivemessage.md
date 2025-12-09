# Model: io.clbs.openhes.models.messaging.MessagingReceiveMessage

Defines a specification for messages delivered from server to a consumer.

## Fields

| Field | Information |
| --- | --- |
| messageId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique message identifier. |
| subject | <b>Type:</b> `string`<br><b>Description:</b><br>The subjects of the message.<br><b>Example:</b> "events.public.acquisition.jobdone" |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The message data. |

