# Model: io.clbs.openhes.models.acquisition.ActionData

Sub-message containing action-based variant of data values

## Fields

| Field | Type | Description |
| --- | --- | --- |
| nodata | google.protobuf.Empty | No data |
| billings | [io.clbs.openhes.models.acquisition.BillingValues](model-io-clbs-openhes-models-acquisition-billingvalues.md) | Register values |
| profile | [io.clbs.openhes.models.acquisition.ProfileValues](model-io-clbs-openhes-models-acquisition-profilevalues.md) | Profile values |
| irregularProfile | [io.clbs.openhes.models.acquisition.IrregularProfileValues](model-io-clbs-openhes-models-acquisition-irregularprofilevalues.md) | Irregular (non-periodical) profile values, e.g. daily profile |
| deviceInfo | [io.clbs.openhes.models.acquisition.DeviceInfo](model-io-clbs-openhes-models-acquisition-deviceinfo.md) | Device info |
| events | [io.clbs.openhes.models.acquisition.EventRecords](model-io-clbs-openhes-models-acquisition-eventrecords.md) | Event records |

