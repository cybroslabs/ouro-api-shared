# Model: io.clbs.openhes.models.messaging.MessagingComponentSpec

Defines the configuration specification for a messaging component.
 Components represent microservices or system modules that publish and/or consume messages.

## Fields

| Field | Information |
| --- | --- |
| enabled | <b>Type:</b> `bool`<br><b>Description:</b><br>When false, the component cannot publish or consume messages. Useful for maintenance or debugging. |
| consumers | <b>Type:</b> [`io.clbs.openhes.models.messaging.MessagingComponentConsumerSettings`](model-io-clbs-openhes-models-messaging-messagingcomponentconsumersettings.md)<br><b>Description:</b><br>The list of consumer configurations for this component. Each consumer handles specific message subjects. |

