# Model: io.clbs.openhes.models.acquisition.JobSettings

Defines the parameters and execution rules that control how a job is processed.

## Fields

| Field | Information |
| --- | --- |
| maxDuration | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The maximum duration allowed for a job attempt. This value defines the real- time window for the driver to execute the job. |
| priority | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobPriority`](enum-io-clbs-openhes-models-acquisition-jobpriority.md)<br><b>Description:</b><br>The execution priority assigned to the job. |
| attempts | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of allowed job attempts. Must be at least `1`. |
| retryDelay | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The delay between two attempts. |
| deferStart | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The time offset to wait before starting the job. |
| expiresAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp that specifies when the job expires and should no longer not be executed. |
| readPathPolicy | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ReadPathPolicy`](enum-io-clbs-openhes-models-acquisition-readpathpolicy.md)<br><b>Description:</b><br>The policy that determines how data is read from devices.  It can be red directly from devices (meters) or via the data concentrator. |

