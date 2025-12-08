# Enum: io.clbs.openhes.models.cronjobs.CronJobTypeEnum

Defines the cron job types that can be scheduled for automated recurring operations.
Cron jobs enable scheduled data collection, configuration synchronization, and other periodic tasks.

## Options

| Value | Description |
| --- | --- |
| CRON_JOB_TYPE_UNSPECIFIED | Unspecified cron job type (invalid, should not be used). |
| CRON_JOB_TYPE_START_BULK | Schedules recurring acquisition bulks for automated data collection from devices. |
| CRON_JOB_TYPE_START_PROXY_BULK | Schedules recurring proxy bulks for data collection operations initiated through the DataProxy. |
