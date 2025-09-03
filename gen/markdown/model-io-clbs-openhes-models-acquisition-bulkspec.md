# Model: io.clbs.openhes.models.acquisition.BulkSpec

Defines a bulk specification.

## Fields

| Field | Information |
| --- | --- |
| correlationId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The correlation identifier that represents a non-homogenous group with various device types. |
| devices | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ListOfJobDeviceId`](model-io-clbs-openhes-models-acquisition-listofjobdeviceid.md)<br><b>Description:</b><br>The list of devices in the bulk. |
| deviceGroupId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique device group identifier. |
| settings | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobSettings`](model-io-clbs-openhes-models-acquisition-jobsettings.md)<br><b>Description:</b><br>The job settings shared across all jobs in the bulk. |
| actions | <b>Type:</b> [`io.clbs.openhes.models.acquisition.JobActionSet`](model-io-clbs-openhes-models-acquisition-jobactionset.md)<br><b>Description:</b><br>The list of actions to be executed on each device in the bulk. |

