# Model: io.clbs.openhes.models.acquisition.GetDeviceBulksRequest

## Fields

| Field | Information |
| --- | --- |
| from | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the bulks selection. If the time range in not specified, the system will return maximum 1000 latest bulks. |
| to | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the bulks selection. If the time range in not specified, the system will return maximum 1000 latest bulks. |
| deviceId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the device. |

