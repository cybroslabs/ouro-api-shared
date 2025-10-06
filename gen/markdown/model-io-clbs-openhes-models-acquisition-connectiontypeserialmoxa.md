# Model: io.clbs.openhes.models.acquisition.ConnectionTypeSerialMoxa

Defines the connection information for a controlled serial line using the Moxa protocol.

## Fields

| Field | Information |
| --- | --- |
| host | <b>Type:</b> `string`<br><b>Description:</b><br>The host name or IP address of the device to connect to. |
| dataPort | <b>Type:</b> `uint32`<br><b>Description:</b><br>The TCP data port number of the device to connect to. |
| commandPort | <b>Type:</b> `uint32`<br><b>Description:</b><br>The TCP command port number of the device to connect to. |
| timeout | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The timeout for the serial port connection using the Moxa protocol. |
| serialConfig | <b>Type:</b> [`io.clbs.openhes.models.acquisition.SerialConfig`](model-io-clbs-openhes-models-acquisition-serialconfig.md)<br><b>Description:</b><br>The serial port configuration. Optional. If not set, current Moxa settings are preserved. |

