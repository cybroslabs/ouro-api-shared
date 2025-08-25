# Model: io.clbs.openhes.models.messaging.MessagingPublisherClient

## Fields

| Field | Information |
| --- | --- |
| setup | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingPublisherSetup`](model-io-clbs-openhes-models-messaging-messagingpublishersetup.md)<br><b>Description:</b><br>Setup action to initialize the publisher. It must be the first message sent defining the publisher. Any other subsequent setup message will be rejected. |
| publish | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingPublishMessage`](model-io-clbs-openhes-models-messaging-messagingpublishmessage.md)<br><b>Description:</b><br>Publish action to send a message to be published. |

