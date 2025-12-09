# Model: io.clbs.openhes.models.messaging.MessagingComponentConsumerSettings

Defines configuration settings for a specific consumer within a messaging component.
 Each consumer represents a logical queue handler that processes messages from subscribed subjects.

## Fields

| Field | Information |
| --- | --- |
| consumerId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique UUID identifier for this consumer instance. Must be unique across all consumers of the same component. |
| maxInFlightMessages | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of unacknowledged messages this consumer can have. Values >1 enable concurrent processing but may lose order. Value 1 guarantees ordered processing.<br><b>Example:</b> 100 |
| subjects | <b>Type:</b> `string`<br><b>Description:</b><br>The message subjects this consumer subscribes to. At least one subject is required. Changes apply to all instances of the component.<br><b>Example:</b> ["events.custom.sapmessage"] |

