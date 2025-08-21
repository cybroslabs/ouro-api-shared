# Model: io.clbs.openhes.models.acquisition.CommunicationUnitLogRecord

Defines a single log record of a communication unit.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the log record. If not provided, a hash of the log data is auto-generated. |
| communicationUnitId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the communication unit to which the log record belongs. May be empty and the relation can be set at a higher object level. |
| timestamp | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp of the log record. Must not be empty. |
| level | <b>Type:</b> [`io.clbs.openhes.models.acquisition.LogRecordLevel`](enum-io-clbs-openhes-models-acquisition-logrecordlevel.md)<br><b>Description:</b><br>The severity level of the log record. |
| message | <b>Type:</b> `string`<br><b>Description:</b><br>The log message. Must not be empty. |

