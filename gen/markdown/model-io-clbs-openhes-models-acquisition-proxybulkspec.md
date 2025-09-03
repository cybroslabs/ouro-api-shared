# Model: io.clbs.openhes.models.acquisition.ProxyBulkSpec

Defines a proxy bulk specification.

## Fields

| Field | Information |
| --- | --- |
| correlationId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The correlation identifier that represents a non-homogenous group with various device types. |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type. |
| devices | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ListOfJobDevice`](model-io-clbs-openhes-models-acquisition-listofjobdevice.md)<br><b>Description:</b><br>The list of devices in the proxy bulk. |
| settings | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobSettings`](model-io-clbs-openhes-models-acquisition-jobsettings.md)<br><b>Description:</b><br>The job settings shared across the bulk. |
| actions | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobAction`](model-io-clbs-openhes-models-acquisition-jobaction.md)<br><b>Description:</b><br>The list of actions to be executed. |

