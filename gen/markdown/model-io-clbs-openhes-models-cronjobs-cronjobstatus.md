# Model: io.clbs.openhes.models.cronjobs.CronJobStatus

Defines the current runtime status of a cron job including execution history and schedule validation.

## Fields

| Field | Information |
| --- | --- |
| lastRunAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the cron job last executed successfully or attempted to execute. |
| nextRunAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The calculated timestamp for the next scheduled execution based on the cron schedule and timezone. |
| error | <b>Type:</b> `string`<br><b>Description:</b><br>An error message if the cron schedule is invalid or if the last execution failed. |

