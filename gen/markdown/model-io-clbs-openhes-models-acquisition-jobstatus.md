# Model: io.clbs.openhes.models.acquisition.JobStatus

Defines the status information of a single job within a bulk.

## Fields

| Field | Information |
| --- | --- |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobStatusCode`](enum-io-clbs-openhes-models-acquisition-jobstatuscode.md)<br><b>Description:</b><br>The status of the job. |
| code | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobErrorCode`](enum-io-clbs-openhes-models-acquisition-joberrorcode.md)<br><b>Description:</b><br>The error code of the job. |
| results | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionResult`](model-io-clbs-openhes-models-acquisition-actionresult.md)<br><b>Description:</b><br>The result data for all executed job actions. |
| startedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the job. |
| finishedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The finish timestamp of the job. |
| attemptsDone | <b>Type:</b> `int32`<br><b>Description:</b><br>The number of attempts already done. |
| deviceInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceInfo`](model-io-clbs-openhes-models-acquisition-deviceinfo.md)<br><b>Description:</b><br>The device informatiom from the `ACTION_TYPE_GET_DEVICE_INFO` action. |
| queueId | <b>Type:</b> `int64`<br><b>Description:</b><br>The internal queue identifier set by the Taskmaster when the job is queued for processing. |
| errorMessage | <b>Type:</b> [`io.clbs.openhes.models.common.FormattedMessage`](model-io-clbs-openhes-models-common-formattedmessage.md)<br><b>Description:</b><br>The user-facing error message related to the whole job. This is used especially relevant when no action was executed, allowing to log non-action-related errors. |

