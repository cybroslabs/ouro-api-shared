# ApiService - Messaging

## CreateMessagingConsumer

Creates a new messaging bi-directional consumer. The stream allows receiving messages and sending acknowledgements.

```proto
CreateMessagingConsumer(io.clbs.openhes.models.messaging.MessagingConsumerClient) returns (io.clbs.openhes.models.messaging.MessagingConsumerServer)
```

- Input: [`io.clbs.openhes.models.messaging.MessagingConsumerClient`](model-io-clbs-openhes-models-messaging-messagingconsumerclient.md)
- Output: [`io.clbs.openhes.models.messaging.MessagingConsumerServer`](model-io-clbs-openhes-models-messaging-messagingconsumerserver.md)

## CreateMessagingPublisher

Creates a new messaging publisher. The stream allows sending messages to be published.

```proto
CreateMessagingPublisher(io.clbs.openhes.models.messaging.MessagingPublisherClient)
```

- Input: [`io.clbs.openhes.models.messaging.MessagingPublisherClient`](model-io-clbs-openhes-models-messaging-messagingpublisherclient.md)

## ListMessagingComponents

Retrieves a paginated list of messaging components based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListMessagingComponents(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.messaging.ListOfMessagingComponent)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.messaging.ListOfMessagingComponent`](model-io-clbs-openhes-models-messaging-listofmessagingcomponent.md)

## UpdateMessagingComponent

Updates the details of an existing messaging component.

```proto
UpdateMessagingComponent(io.clbs.openhes.models.messaging.MessagingComponent)
```

- Input: [`io.clbs.openhes.models.messaging.MessagingComponent`](model-io-clbs-openhes-models-messaging-messagingcomponent.md)

