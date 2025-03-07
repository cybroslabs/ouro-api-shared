# Model: io.clbs.openhes.models.acquisition.ActionData

Sub-message containing action-based variant of data values

## Fields

| Field | Type | Description |
| --- | --- | --- |
| nodata | google.protobuf.Empty | No data |
| billings | io.clbs.openhes.models.acquisition.BillingValues | Register values |
| profile | io.clbs.openhes.models.acquisition.ProfileValues | Profile values |
| irregularProfile | io.clbs.openhes.models.acquisition.IrregularProfileValues | Irregular (non-periodical) profile values, e.g. daily profile |
| deviceInfo | io.clbs.openhes.models.acquisition.DeviceInfo | Device info |
| events | io.clbs.openhes.models.acquisition.EventRecords | Event records |

