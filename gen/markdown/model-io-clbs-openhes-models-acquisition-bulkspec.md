# Model: io.clbs.openhes.models.acquisition.BulkSpec

Defines a bulk specification containing the configuration for a mass operation.

## Fields

| Field | Information |
| --- | --- |
| correlationId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>A correlation identifier for grouping related bulks across different driver types. |
| devices | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ListOfJobDeviceId`](model-io-clbs-openhes-models-acquisition-listofjobdeviceid.md)<br><b>Description:</b><br>An explicit list of device identifiers to target. |
| deviceGroupId | <b>Type:</b> `string`<br><b>Description:</b><br>A device group identifier to target all devices in the group. |
| settings | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobSettings`](model-io-clbs-openhes-models-acquisition-jobsettings.md)<br><b>Description:</b><br>Execution settings including priority, retry logic, and timeout values. |
| actions | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionSet`](model-io-clbs-openhes-models-acquisition-jobactionset.md)<br><b>Description:</b><br>The sequence of actions to execute on each device (e.g., read registers, update firmware). |
| skipDataProcessing | <b>Type:</b> `bool`<br><b>Description:</b><br>When true, collected data bypasses normal processing pipelines (useful for testing). Default is false. |

