# Model: io.clbs.openhes.models.cronjobs.CronJobStatus

Defines the current status of a cron job.

## Fields

| Field | Information |
| --- | --- |
| lastRunAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp of the last run of the cron job. |
| nextRunAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp of the next run of the cron job. |
| error | <b>Type:</b> `string`<br><b>Description:</b><br>The error message if the cron job schedule is invalid. |

