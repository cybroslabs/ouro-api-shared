# Model: io.clbs.openhes.models.cronjobs.CronJobSpec

Defines the cron job specification containing scheduling and execution details.

## Fields

| Field | Information |
| --- | --- |
| type | <b>Type:</b> [`io.clbs.openhes.models.cronjobs.CronJobTypeEnum`](enum-io-clbs-openhes-models-cronjobs-cronjobtypeenum.md)<br><b>Description:</b><br>The type of operation to execute (e.g., START_BULK, START_PROXY_BULK). |
| schedule | <b>Type:</b> `string`<br><b>Description:</b><br>The cron schedule expression in standard format (e.g., "0 * * * *" for hourly, "0 0 * * *" for daily at midnight). |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone for schedule interpretation using IANA time zone database names (e.g., "America/New_York", "Europe/Prague", "UTC"). |
| suspend | <b>Type:</b> `bool`<br><b>Description:</b><br>When true, the cron job is paused and will not execute until resumed. Default is false. |
| data | <b>Type:</b> `google.protobuf.Struct`<br><b>Description:</b><br>The type-specific payload passed to the operation when it executes (e.g., bulk specification for START_BULK type). |

