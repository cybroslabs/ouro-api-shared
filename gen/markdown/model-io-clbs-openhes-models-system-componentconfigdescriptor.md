# Model: io.clbs.openhes.models.system.ComponentConfigDescriptor

Defines the configuration schema for a component, declaring what fields are available.
 Components register their configuration schema at startup to enable dynamic configuration.

## Fields

| Field | Information |
| --- | --- |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The unique component name. |
| items | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)<br><b>Description:</b><br>The field descriptors defining available configuration options including types, defaults, and validation rules. |

