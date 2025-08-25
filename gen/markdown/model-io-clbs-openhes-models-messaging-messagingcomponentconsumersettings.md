# Model: io.clbs.openhes.models.messaging.MessagingComponentConsumerSettings

## Fields

| Field | Information |
| --- | --- |
| consumerId | <b>Type:</b> `string`<br><b>Description:</b><br>Unique UUID identifier for the consumer. It must be unique across all consumers of the same component. |
| deliveryPolicy | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingDeliveryPolicy`](enum-io-clbs-openhes-models-messaging-messagingdeliverypolicy.md)<br><b>Description:</b><br>Delivery policy for the consumer. |
| maxInFlightMessages | <b>Type:</b> `int32`<br><b>Description:</b><br>Maximum number of in-flight (unacknowledged) messages allowed for the consumer. Value >1 means that multiple consumers may receive messages concurrently without preserving order. Value 1 means that only a single message may be in-flight at any time, preserving order even when multiple consumers are running. |
| subjects | <b>Type:</b> `string`<br><b>Description:</b><br>List of subjects the consumer is interested in. It must not be empty. If it's changed then all existing consumers for the same component will be affected! |

