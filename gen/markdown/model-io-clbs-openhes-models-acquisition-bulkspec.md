# Model: io.clbs.openhes.models.acquisition.BulkSpec

## Fields

| Field | Information |
| --- | --- |
| correlationId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br><br>The correlation identifier, e.g. to define relation to non-homogenous group. |
| devices | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ListOfJobDeviceId`](model-io-clbs-openhes-models-acquisition-listofjobdeviceid.md)<br><b>Description:</b><br>The list of devices in the bulk. |
| deviceGroupId | <b>Type:</b> `string`<br><b>Description:</b><br>The device group identifier. |
| settings | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobSettings`](model-io-clbs-openhes-models-acquisition-jobsettings.md)<br><b>Description:</b><br>The bulk-shared job settings. |
| actions | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionSet`](model-io-clbs-openhes-models-acquisition-jobactionset.md)<br><b>Description:</b><br>The list actions to be executed. |
| webhookUrl | <b>Type:</b> `string`<br><b>Description:</b><br>The webhook URL to call when the bulk is completed. |

