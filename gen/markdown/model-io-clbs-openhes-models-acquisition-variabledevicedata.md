# Model: io.clbs.openhes.models.acquisition.VariableDeviceData

Defines data for a specific variable within a device.

## Fields

| Field | Information |
| --- | --- |
| variableId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique variable identifier. |
| timestamps | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>A list of timestamps for the variable data. |
| units | <b>Type:</b> `string`<br><b>Description:</b><br>A list of units for the variable data. |
| values | <b>Type:</b> [`io.clbs.openhes.models.acquisition.MeasuredValue`](model-io-clbs-openhes-models-acquisition-measuredvalue.md)<br><b>Description:</b><br>A list of measured values for the variable data. |

