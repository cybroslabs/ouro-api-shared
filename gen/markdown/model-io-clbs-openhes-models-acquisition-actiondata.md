# Model: io.clbs.openhes.models.acquisition.ActionData

Defines the action-based variant of data values.

## Fields

| Field | Information |
| --- | --- |
| nodata | <b>Type:</b> `google.protobuf.Empty`<br><b>Description:</b><br>Indicates that no data was returned by the action. |
| registers | <b>Type:</b> [`io.clbs.openhes.models.acquisition.RegisterValues`](model-io-clbs-openhes-models-acquisition-registervalues.md)<br><b>Description:</b><br>The set of register values. |
| profile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ProfileValues`](model-io-clbs-openhes-models-acquisition-profilevalues.md)<br><b>Description:</b><br>The set of profile values. |
| irregularProfile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.IrregularProfileValues`](model-io-clbs-openhes-models-acquisition-irregularprofilevalues.md)<br><b>Description:</b><br>The set of irregular (non-periodical) profile values, such as daily profiles. |
| deviceInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceInfo`](model-io-clbs-openhes-models-acquisition-deviceinfo.md)<br><b>Description:</b><br>The device information. |
| events | <b>Type:</b> [`io.clbs.openhes.models.acquisition.EventRecords`](model-io-clbs-openhes-models-acquisition-eventrecords.md)<br><b>Description:</b><br>The event records. |
| touTable | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.TimeOfUseTableSpec`](model-io-clbs-openhes-models-acquisition-timeofuse-timeofusetablespec.md)<br><b>Description:</b><br>The time-of-use (TOU) table. |

