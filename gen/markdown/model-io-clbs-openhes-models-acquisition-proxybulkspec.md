# Model: io.clbs.openhes.models.acquisition.ProxyBulkSpec

Defines a proxy bulk specification for operations forwarded from the DataProxy to the main API.
 Proxy bulks are used when data collection is initiated externally and needs to be tracked in the main system.

## Fields

| Field | Information |
| --- | --- |
| correlationId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>A correlation identifier linking this proxy bulk to the originating bulk in the DataProxy. |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type that will execute the jobs (e.g., "dlms", "mbus"). |
| devices | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ListOfJobDevice`](model-io-clbs-openhes-models-acquisition-listofjobdevice.md)<br><b>Description:</b><br>The list of devices with their connection information for the proxy bulk. |
| settings | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobSettings`](model-io-clbs-openhes-models-acquisition-jobsettings.md)<br><b>Description:</b><br>Execution settings including priority, retry logic, and timeout values. |
| actions | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobAction`](model-io-clbs-openhes-models-acquisition-jobaction.md)<br><b>Description:</b><br>The sequence of driver-specific actions to execute on each device. |

