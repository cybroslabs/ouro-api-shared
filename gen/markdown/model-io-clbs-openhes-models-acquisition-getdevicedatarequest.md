# Model: io.clbs.openhes.models.acquisition.GetDeviceDataRequest

Defines a specification for retrieving device data.

## Fields

| Field | Information |
| --- | --- |
| rangeStart | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the requested data. The exclusivness/inclusiveness of the timestamp is defined in a variable. |
| rangeEnd | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the data. The timestamp is inclusive. |
| series | <b>Type:</b> [`io.clbs.openhes.models.acquisition.GetDeviceDataSeriesSelector`](model-io-clbs-openhes-models-acquisition-getdevicedataseriesselector.md)<br><b>Description:</b><br>One or more selectors identifying devices and variables for which data is requested. |
| filterIncludeStatus | <b>Type:</b> `int64`<br><b>Description:</b><br>The status bit filter. Only values with these status bits will be included in the response. |
| filterExcludeStatus | <b>Type:</b> `int64`<br><b>Description:</b><br>The status bit filter. Values with these status bits will not be included in the response. |
| snapshot | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>If set, returns a snapshot of stoered data from the given point in time. |

