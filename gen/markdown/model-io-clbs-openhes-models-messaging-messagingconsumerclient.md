# Model: io.clbs.openhes.models.messaging.MessagingConsumerClient

Defines the specification for consumer messages.

## Fields

| Field | Information |
| --- | --- |
| setup | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingConsumerSetup`](model-io-clbs-openhes-models-messaging-messagingconsumersetup.md)<br><b>Description:</b><br>Setup action to initialize the consumer. Must be the first message sent. Any subsequent setup messages are rejected. |
| ack | <b>Type:</b> `google.protobuf.StringValue`<br><b>Description:</b><br>Acknowledgement action to confirm the message with the given message ID. |
| nak | <b>Type:</b> `google.protobuf.StringValue`<br><b>Description:</b><br>Negative acknowledgement action to reject and requeue the message with the given message ID. |

