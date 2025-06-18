# Enum: io.clbs.openhes.models.acquisition.JobErrorCode

Error codes related to jobs

## Options

| Value | Description |
| --- | --- |
| JOB_ERROR_CODE_NONE | The job has been completed successfully. |
| JOB_ERROR_CODE_BUSY | There is no free slot in the driver to handle the job; the job shall be send again later. |
| JOB_ERROR_CODE_ERROR | The job failed, the retry will be attempted. |
| JOB_ERROR_CODE_ALREADY_EXISTS | This should never happen! It means that the same job is currently being processed by the driver and it was sent to the driver mutliple times which would mean that there is a bug. |
| JOB_ERROR_CODE_FATAL | The job failed, the retry will NOT be attempted. |
