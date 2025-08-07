# Model: io.clbs.openhes.models.acquisition.CommunicationUnitLogRecord

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the log record, if provided. If not provided, a hash of the log data is auto-generated. |
| communicationUnitId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the communication unit to which the log record belongs. When adding the log records, it may be empty and relation can be set on the upper object level instead. |
| timestamp | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp of the log record. It must not be empty. |
| level | <b>Type:</b> [`io.clbs.openhes.models.acquisition.LogRecordLevel`](enum-io-clbs-openhes-models-acquisition-logrecordlevel.md)<br><b>Description:</b><br>The log level of the log record, if provided. |
| message | <b>Type:</b> `string`<br><b>Description:</b><br>The log message. It must not be empty. |

