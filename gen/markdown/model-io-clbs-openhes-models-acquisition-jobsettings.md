# Model: io.clbs.openhes.models.acquisition.JobSettings

Sub-message containing job parameters

## Fields

| Field | Type | Description |
| --- | --- | --- |
| maxDuration | google.protobuf.Duration | Maximum duration of the job attempt. This is related to the real time for the driver. |
| priority | [io.clbs.openhes.models.acquisition.JobPriority](model-io-clbs-openhes-models-acquisition-jobpriority.md) | Priority of the job. |
| attempts | int32 | Maximum number of attempts, 1 is the minimum. |
| retryDelay | google.protobuf.Duration | Delay between two attempts. |
| deferStart | google.protobuf.Duration | Time to wait before starting the job. |
| expiresAt | google.protobuf.Timestamp | The timestamp when the job expires. |

