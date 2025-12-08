# Model: io.clbs.openhes.models.acquisition.JobSettings

Defines the parameters and execution rules that control how a job is processed.
 These settings apply to individual jobs or entire bulks and determine retry logic, timeouts, and scheduling.

## Fields

| Field | Information |
| --- | --- |
| maxDuration | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The maximum real-time duration allowed for a single job attempt before it's terminated. |
| priority | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobPriority`](enum-io-clbs-openhes-models-acquisition-jobpriority.md)<br><b>Description:</b><br>The execution priority determining queue position relative to other jobs. |
| attempts | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of retry attempts if the job fails. Must be at least 1. Each attempt index can have different settings. |
| retryDelay | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The delay to wait between retry attempts after a failure. |
| deferStart | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The initial delay before the job becomes eligible for execution (useful for scheduling). |
| expiresAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The expiration timestamp after which the job will not be executed, even if still queued. |
| readPathPolicy | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ReadPathPolicy`](enum-io-clbs-openhes-models-acquisition-readpathpolicy.md)<br><b>Description:</b><br>The strategy for reading data: directly from meters, via data concentrator, or automatic fallback. |

