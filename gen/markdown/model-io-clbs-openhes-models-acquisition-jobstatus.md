# Model: io.clbs.openhes.models.acquisition.JobStatus

Sub-message containing job status info

## Fields

| Field | Information |
| --- | --- |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobStatusCode`](model-io-clbs-openhes-models-acquisition-jobstatuscode.md)<br><b>Description:</b><br>The status of the job. |
| code | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobErrorCode`](model-io-clbs-openhes-models-acquisition-joberrorcode.md)<br><b>Description:</b><br>The error code of the job. |
| results | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionResult`](model-io-clbs-openhes-models-acquisition-actionresult.md)<br><b>Description:</b><br>The result data for all job actions. |
| startedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the job. |
| finishedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The finish timestamp of the job. |
| attemptsDone | <b>Type:</b> `int32`<br><b>Description:</b><br>The number of attempts already done. |
| deviceInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceInfo`](model-io-clbs-openhes-models-acquisition-deviceinfo.md)<br><b>Description:</b><br>The device info. It contains the data from the action ACTION_TYPE_GET_DEVICE_INFO. |
| queueId | <b>Type:</b> `int64`<br><b>Description:</b><br>The internal queue identifier set by the Taskmaster when the job is queued to process. |

