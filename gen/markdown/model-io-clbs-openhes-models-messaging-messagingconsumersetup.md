# Model: io.clbs.openhes.models.messaging.MessagingConsumerSetup

## Fields

| Field | Information |
| --- | --- |
| componentId | <b>Type:</b> `string`<br><b>Description:</b><br>Component ID to which the consumer is associated. |
| deliveryPolicy | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingDeliveryPolicy`](enum-io-clbs-openhes-models-messaging-messagingdeliverypolicy.md)<br><b>Description:</b><br>Delivery policy for the consumer. |
| maxInFlightMessages | <b>Type:</b> `int32`<br><b>Description:</b><br>Maximum number of in-flight (unacknowledged) messages allowed for the consumer. Value >1 means that multiple consumers may receive messages concurrently without preserving order. Value 1 means that only a single message may be in-flight at any time, preserving order even when multiple consumers are running. |

