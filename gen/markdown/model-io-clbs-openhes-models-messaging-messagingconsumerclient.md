# Model: io.clbs.openhes.models.messaging.MessagingConsumerClient

Defines messages sent from a consumer client to the messaging server in a bidirectional stream.
 Consumers use this to establish a connection, acknowledge processed messages, or reject failed messages.

## Fields

| Field | Information |
| --- | --- |
| setup | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingConsumerSetup`](model-io-clbs-openhes-models-messaging-messagingconsumersetup.md)<br><b>Description:</b><br>Setup message to initialize the consumer connection. Must be the first message sent. Subsequent setup messages are rejected. |
| ack | <b>Type:</b> `google.protobuf.StringValue`<br><b>Description:</b><br>Acknowledgement confirming successful processing of a message. The message is removed from the queue. |
| nak | <b>Type:</b> `google.protobuf.StringValue`<br><b>Description:</b><br>Negative acknowledgement indicating processing failure. The message is requeued for redelivery. |

