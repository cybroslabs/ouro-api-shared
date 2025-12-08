# Model: io.clbs.openhes.models.messaging.MessagingConsumerServer

Defines messages sent from the messaging server to a consumer client in a bidirectional stream.
 The server delivers messages that match the consumer's subscribed subjects.

## Fields

| Field | Information |
| --- | --- |
| receive | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingReceiveMessage`](model-io-clbs-openhes-models-messaging-messagingreceivemessage.md)<br><b>Description:</b><br>Delivers a message to the consumer for processing. The consumer must ack or nak this message. |

