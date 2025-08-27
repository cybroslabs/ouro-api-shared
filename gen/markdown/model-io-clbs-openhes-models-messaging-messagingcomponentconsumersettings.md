# Model: io.clbs.openhes.models.messaging.MessagingComponentConsumerSettings

## Fields

| Field | Information |
| --- | --- |
| consumerId | <b>Type:</b> `string`<br><b>Description:</b><br>Unique UUID identifier for the consumer. It must be unique across all consumers of the same component. |
| maxInFlightMessages | <b>Type:</b> `int32`<br><b>Description:</b><br>Maximum number of in-flight (unacknowledged) messages allowed for the consumer. Value >1 means that multiple consumers may receive messages concurrently without preserving order. Value 1 means that only a single message may be in-flight at any time, preserving order even when multiple consumers are running. |
| subjects | <b>Type:</b> `string`<br><b>Description:</b><br>List of non-empty subjects the consumer is interested in. At least one must be set. If it's changed then all instances of the same component will be affected! |

