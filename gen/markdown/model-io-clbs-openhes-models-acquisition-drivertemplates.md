# Model: io.clbs.openhes.models.acquisition.DriverTemplates

Defines the driver templates exchanged during driver negotiation.

## Fields

| Field | Information |
| --- | --- |
| communicationTemplates | <b>Type:</b> [`io.clbs.openhes.models.acquisition.CommunicationTemplate`](model-io-clbs-openhes-models-acquisition-communicationtemplate.md)<br><b>Description:</b><br>The supported communication options templates. Each template represents one communication method. A driver can support multiple communication templates. |
| appProtocols | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocolTemplate`](model-io-clbs-openhes-models-acquisition-applicationprotocoltemplate.md)<br><b>Description:</b><br>The supported application protocol templates. |
| actionAttributes | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionAttributes`](model-io-clbs-openhes-models-acquisition-jobactionattributes.md)<br><b>Description:</b><br>The job action templates for all supported action types. Each supported action type must appear only once. |
| accessTemplates | <b>Type:</b> [`io.clbs.openhes.models.acquisition.AccessLevelTemplate`](model-io-clbs-openhes-models-acquisition-accessleveltemplate.md)<br><b>Description:</b><br>The supported access level templates. |
| actionConstraints | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionContraints`](model-io-clbs-openhes-models-acquisition-jobactioncontraints.md)<br><b>Description:</b><br>The supported templates of the job action constraints. |
| uknownDeviceDescriptors | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)<br><b>Description:</b><br>The list of field descriptors for unknown devices detected by the communication unit.<br>This applies only to drivers that communicate with devices like data concentrators that can provide information for unknown devices.<br>The descriptors must cover all data attributes used in the `SetUnknownDevicesRequest` message. |

