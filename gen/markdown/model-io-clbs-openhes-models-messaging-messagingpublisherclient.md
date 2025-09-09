# Model: io.clbs.openhes.models.messaging.MessagingPublisherClient

Defines the specification for published messages.

## Fields

| Field | Information |
| --- | --- |
| setup | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingPublisherSetup`](model-io-clbs-openhes-models-messaging-messagingpublishersetup.md)<br><b>Description:</b><br>Setup action to initialize the publisher. Must be the first message sent. Any subsequent setup messages are rejected. |
| publish | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingPublishMessage`](model-io-clbs-openhes-models-messaging-messagingpublishmessage.md)<br><b>Description:</b><br>Publish action to send a message. |

