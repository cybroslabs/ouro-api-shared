# Model: io.clbs.openhes.models.acquisition.ActionData

Sub-message containing action-based variant of data values

## Fields

| Field | Information |
| --- | --- |
| nodata | <b>Type:</b> `google.protobuf.Empty`<br><b>Description:</b><br>No data |
| registers | <b>Type:</b> [`io.clbs.openhes.models.acquisition.RegisterValues`](model-io-clbs-openhes-models-acquisition-registervalues.md)<br><b>Description:</b><br>Register values |
| profile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ProfileValues`](model-io-clbs-openhes-models-acquisition-profilevalues.md)<br><b>Description:</b><br>Profile values |
| irregularProfile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.IrregularProfileValues`](model-io-clbs-openhes-models-acquisition-irregularprofilevalues.md)<br><b>Description:</b><br>Irregular (non-periodical) profile values, e.g. daily profile |
| deviceInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceInfo`](model-io-clbs-openhes-models-acquisition-deviceinfo.md)<br><b>Description:</b><br>Device info |
| events | <b>Type:</b> [`io.clbs.openhes.models.acquisition.EventRecords`](model-io-clbs-openhes-models-acquisition-eventrecords.md)<br><b>Description:</b><br>Event records |
| touTable | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.TimeOfUseTableSpec`](model-io-clbs-openhes-models-acquisition-timeofuse-timeofusetablespec.md)<br><b>Description:</b><br>The time-of-use table. |

