# Model: io.clbs.openhes.models.system.ComponentConfig

Defines the configuration for a single system component or microservice.
 Each component can define its own configuration fields through field descriptors.

## Fields

| Field | Information |
| --- | --- |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The unique component name (e.g., "taskmaster", "dataproxy", "api"). |
| items | <b>Type:</b> [`io.clbs.openhes.models.common.FieldValues`](model-io-clbs-openhes-models-common-fieldvalues.md)<br><b>Description:</b><br>The configuration values as field-value pairs. Fields are defined by the component's descriptors. |

