# Enum: io.clbs.openhes.models.acquisition.JobErrorCode

Defines the error codes related to jobs.

## Options

| Value | Description |
| --- | --- |
| JOB_ERROR_CODE_UNSPECIFIED | Unspecified job error code. |
| JOB_ERROR_CODE_NONE | The job has been completed successfully. |
| JOB_ERROR_CODE_BUSY | There is no free slot in the driver to handle the job; the job shall be sent again later. |
| JOB_ERROR_CODE_ERROR | The job has failed; a retry will be attempted. |
| JOB_ERROR_CODE_ALREADY_EXISTS | This should never happen! It indicates that the same job is currently being processed by the driver and was sent multiple times, which would point to a bug. |
| JOB_ERROR_CODE_FATAL | The job failed; no retry will be attempted. |
