# Model: io.clbs.openhes.models.acquisition.ProxyBulkSpec

## Fields

| Field | Type | Description |
| --- | --- | --- |
| correlationId | string | @gqltype: UUID<br><br>The correlation identifier, e.g. to define relation to non-homogenous group. |
| driverType | string | The device (driver) type. |
| devices | io.clbs.openhes.models.acquisition.ListOfJobDevice | The list of custom devices in the bulk. |
| settings | io.clbs.openhes.models.acquisition.JobSettings | The bulk-shared job settings. |
| actions | io.clbs.openhes.models.acquisition.JobAction | The list actions to be executed. |
| webhookUrl | string | The webhook URL to call when the bulk is completed. |

