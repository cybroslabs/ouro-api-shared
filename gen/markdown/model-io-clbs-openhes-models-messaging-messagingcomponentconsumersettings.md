# Model: io.clbs.openhes.models.messaging.MessagingComponentConsumerSettings

Defines a specification of messaging component settings for a specified consumer.

## Fields

| Field | Information |
| --- | --- |
| consumerId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique UUID identifier for the consumer. Must be unique across all consumers of the same component. |
| maxInFlightMessages | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of in-flight (unacknowledged) messages allowed. A value `>1` allows multiple consumers to receive messages concurrently without preserving order. Value `1` allows only a single in-flight message at any time and preserves the message order. |
| subjects | <b>Type:</b> `string`<br><b>Description:</b><br>The list of non-empty subjects the consumer is interested in. At least one subject must be set. Changes affect all instances of the same component! |

