# Model: io.clbs.openhes.models.cronjobs.CronJob

Defines the cron job model representing a scheduled recurring task.
 Cron jobs are persistent scheduled operations that execute automatically according to their schedule definition.

## Fields

| Field | Information |
| --- | --- |
| spec | <b>Type:</b> [`io.clbs.openhes.models.cronjobs.CronJobSpec`](model-io-clbs-openhes-models-cronjobs-cronjobspec.md)<br><b>Description:</b><br>The configuration defining when and what to execute. |
| status | <b>Type:</b> [`io.clbs.openhes.models.cronjobs.CronJobStatus`](model-io-clbs-openhes-models-cronjobs-cronjobstatus.md)<br><b>Description:</b><br>The runtime status including last/next execution times and errors. |
| metadata | <b>Type:</b> [`io.clbs.openhes.models.common.MetadataFields`](model-io-clbs-openhes-models-common-metadatafields.md)<br><b>Description:</b><br>Metadata including id, name, generation, user-managed fields, and system-managed fields. |

