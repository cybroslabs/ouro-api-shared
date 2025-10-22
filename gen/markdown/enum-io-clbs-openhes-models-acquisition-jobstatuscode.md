# Enum: io.clbs.openhes.models.acquisition.JobStatusCode

Defines the possible statuses of jobs.

## Options

| Value | Description |
| --- | --- |
| JOB_STATUS_UNSPECIFIED | Unspecified job status. |
| JOB_STATUS_QUEUED | The job is waiting in the queue for execution. |
| JOB_STATUS_RUNNING | The job is currently running. |
| JOB_STATUS_PROCESSING_DATA | The job has finished acquiring data and is now processing it. |
| JOB_STATUS_COMPLETED | The job has been completed. |
| JOB_STATUS_FAILED | The job has failed. |
| JOB_STATUS_CANCELLING | The job is being cancelled. This is a transient state when the parent bulk or the job itself has been requested to be cancelled. When the cancellation is completed, the job status changes to `CANCELLED`. |
| JOB_STATUS_CANCELLED | The job has been cancelled. |
| JOB_STATUS_EXPIRED | The job has expired. |
