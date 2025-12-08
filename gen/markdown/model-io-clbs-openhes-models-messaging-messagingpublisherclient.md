# Model: io.clbs.openhes.models.messaging.MessagingPublisherClient

Defines messages sent from a publisher client to the messaging server in a client stream.
 Publishers use this to establish a connection and send messages to specific subjects.

## Fields

| Field | Information |
| --- | --- |
| setup | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingPublisherSetup`](model-io-clbs-openhes-models-messaging-messagingpublishersetup.md)<br><b>Description:</b><br>Setup message to initialize the publisher connection. Must be the first message sent. Subsequent setup messages are rejected. |
| publish | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingPublishMessage`](model-io-clbs-openhes-models-messaging-messagingpublishmessage.md)<br><b>Description:</b><br>Publishes a message to a subject. All consumers subscribed to this subject will receive it. |

