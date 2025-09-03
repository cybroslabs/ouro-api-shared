# Model: io.clbs.openhes.models.acquisition.BulkStatus

Defines the status of a bulk.

## Fields

| Field | Information |
| --- | --- |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.BulkStatusCode`](enum-io-clbs-openhes-models-acquisition-bulkstatuscode.md)<br><b>Description:</b><br>The overall bulk status. |
| jobsCount | <b>Type:</b> `int32`<br><b>Description:</b><br>The total number of jobs in the bulk. |
| jobsFinished | <b>Type:</b> `int32`<br><b>Description:</b><br>The number of jobs that have finished. |
| jobsSuccessful | <b>Type:</b> `int32`<br><b>Description:</b><br>The number of jobs that finished successfully. |
| createdAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the bulk was created. |
| startedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the bulk started. |
| finishedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the bulk finished. |

