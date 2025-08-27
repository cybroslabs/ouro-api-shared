# Model: io.clbs.openhes.models.acquisition.ProxyBulkSpec

## Fields

| Field | Information |
| --- | --- |
| correlationId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The correlation identifier, e.g. to define relation to non-homogenous group. |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The device (driver) type. |
| devices | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ListOfJobDevice`](model-io-clbs-openhes-models-acquisition-listofjobdevice.md)<br><b>Description:</b><br>The list of custom devices in the proxy bulk. |
| settings | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobSettings`](model-io-clbs-openhes-models-acquisition-jobsettings.md)<br><b>Description:</b><br>The bulk-shared job settings. |
| actions | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobAction`](model-io-clbs-openhes-models-acquisition-jobaction.md)<br><b>Description:</b><br>The list actions to be executed. |

