# Model: io.clbs.openhes.models.cronjobs.CronJobSpec

## Fields

| Field | Information |
| --- | --- |
| type | <b>Type:</b> [`io.clbs.openhes.models.cronjobs.CronJobTypeEnum`](enum-io-clbs-openhes-models-cronjobs-cronjobtypeenum.md)<br><b>Description:</b><br>The type of the cron job, e.g., CRON_JOB_TYPE_START_BULK, CRON_JOB_TYPE_START_PROXY_BULK. |
| schedule | <b>Type:</b> `string`<br><b>Description:</b><br>The cron job definition, e.g., "0 * * * *" for every hour. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone related to the cron job, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2". |
| suspend | <b>Type:</b> `bool`<br><b>Description:</b><br>Whether the cron job is suspended or not. By default, it is false (i.e., not suspended). |
| data | <b>Type:</b> `google.protobuf.Struct`<br><b>Description:</b><br>The payload to be sent when the cron job runs. |

