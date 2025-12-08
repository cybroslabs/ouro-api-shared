# Enum: io.clbs.openhes.models.acquisition.JobPriority

Defines the available job priority levels for controlling execution order.
Higher priority jobs are executed before lower priority jobs when resources are constrained.
Use higher priorities for time-sensitive operations like remote disconnect/connect commands.

## Options

| Value | Description |
| --- | --- |
| JOB_PRIORITY_UNSPECIFIED | Unspecified priority (treated as default/medium priority). |
| JOB_PRIORITY_0 | Lowest priority - for bulk data collection and non-urgent tasks. |
| JOB_PRIORITY_1 | Very low priority. |
| JOB_PRIORITY_2 | Low priority. |
| JOB_PRIORITY_3 | Below-medium priority. |
| JOB_PRIORITY_4 | Medium priority - default for regular data collection. |
| JOB_PRIORITY_5 | Above-medium priority. |
| JOB_PRIORITY_6 | High priority - for important configuration changes. |
| JOB_PRIORITY_7 | Very high priority - for urgent operations. |
| JOB_PRIORITY_8 | Highest priority - for critical operations like emergency disconnects. |
