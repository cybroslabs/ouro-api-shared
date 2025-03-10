# Model: io.clbs.openhes.models.acquisition.DriverTemplates

Sub-message in driver negotiation request

## Fields

| Field | Type | Description |
| --- | --- | --- |
| communicationTemplates | [io.clbs.openhes.models.acquisition.CommunicationTemplate](model-io-clbs-openhes-models-acquisition-communicationtemplate.md) | The templates of the communication options. Every template represents an option how the drivers allows to communicate. The driver can support multiple communication templates. |
| appProtocols | [io.clbs.openhes.models.acquisition.ApplicationProtocolTemplate](model-io-clbs-openhes-models-acquisition-applicationprotocoltemplate.md) | The templates of the application protocols supported by the driver. |
| actionAttributes | [io.clbs.openhes.models.acquisition.JobActionAttributes](model-io-clbs-openhes-models-acquisition-jobactionattributes.md) | The templates of the job actions for all supported action types. It must contain every action type supported by the driver once and only once. |
| accessTemplates | [io.clbs.openhes.models.acquisition.AccessLevelTemplate](model-io-clbs-openhes-models-acquisition-accessleveltemplate.md) | The templates of the access levels supported by the driver. |
| actionConstraints | [io.clbs.openhes.models.acquisition.JobActionContraints](model-io-clbs-openhes-models-acquisition-jobactioncontraints.md) | The templates of the job actions constraints. |

