# Model: io.clbs.openhes.models.acquisition.DriverTemplates

Sub-message in driver negotiation request

## Fields

| Field | Information |
| --- | --- |
| communicationTemplates | <b>Type:</b> [`io.clbs.openhes.models.acquisition.CommunicationTemplate`](model-io-clbs-openhes-models-acquisition-communicationtemplate.md)<br><b>Description:</b><br>The templates of the communication options. Every template represents an option how the drivers allows to communicate. The driver can support multiple communication templates. |
| appProtocols | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocolTemplate`](model-io-clbs-openhes-models-acquisition-applicationprotocoltemplate.md)<br><b>Description:</b><br>The templates of the application protocols supported by the driver. |
| actionAttributes | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionAttributes`](model-io-clbs-openhes-models-acquisition-jobactionattributes.md)<br><b>Description:</b><br>The templates of the job actions for all supported action types. It must contain every action type supported by the driver once and only once. |
| accessTemplates | <b>Type:</b> [`io.clbs.openhes.models.acquisition.AccessLevelTemplate`](model-io-clbs-openhes-models-acquisition-accessleveltemplate.md)<br><b>Description:</b><br>The templates of the access levels supported by the driver. |
| actionConstraints | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionContraints`](model-io-clbs-openhes-models-acquisition-jobactioncontraints.md)<br><b>Description:</b><br>The templates of the job actions constraints. |
| uknownDeviceDescriptors | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)<br><b>Description:</b><br>The list of descriptors related to the uknown devices visible to the communication unit.<br> This relates only to the drivers that communications with a device like data-concentrator that can provide information for unknown devices.<br> The descripts must cover all data attributes which drivers sets in the SetUnknownDevicesRequest message. |

