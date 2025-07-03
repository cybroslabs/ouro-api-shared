# Model: io.clbs.openhes.models.acquisition.GetDeviceDataRequest

## Fields

| Field | Information |
| --- | --- |
| from | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the data. |
| to | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the data. |
| series | <b>Type:</b> [`io.clbs.openhes.models.acquisition.GetDeviceDataSeriesSelector`](model-io-clbs-openhes-models-acquisition-getdevicedataseriesselector.md)<br><b>Description:</b><br>One or more identifiers of the devices. |
| filterIncludeStatus | <b>Type:</b> `int64`<br><b>Description:</b><br>The filter status bits, only values with these status bits will be returned in the response. |
| filterExcludeStatus | <b>Type:</b> `int64`<br><b>Description:</b><br>The filter status bits, only values without these status bits will be returned in the response. |
| snapshot | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>If set, the system will return a snapshot from the given point in time, e.g. what has been stored in the system yesterday. |

