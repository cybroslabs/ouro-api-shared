# Model: io.clbs.openhes.models.acquisition.JobSettings

Sub-message containing job parameters

## Fields

| Field | Information |
| --- | --- |
| maxDuration | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>Maximum duration of the job attempt. This is related to the real time for the driver. |
| priority | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobPriority`](enum-io-clbs-openhes-models-acquisition-jobpriority.md)<br><b>Description:</b><br>Priority of the job. |
| attempts | <b>Type:</b> `int32`<br><b>Description:</b><br>Maximum number of attempts, 1 is the minimum. |
| retryDelay | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>Delay between two attempts. |
| deferStart | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>Time to wait before starting the job. |
| expiresAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the job expires. |

