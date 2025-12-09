# Model: io.clbs.openhes.models.acquisition.GetDeviceBulksRequest

Defines a specification for retrieving device bulks from a specified time range.

## Fields

| Field | Information |
| --- | --- |
| rangeStart | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the bulks selection. It is inclusive. |
| rangeEnd | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the bulks selection. It is inclusive. |
| deviceId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique identifier of the device. |

