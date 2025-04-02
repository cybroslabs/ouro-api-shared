# Model: io.clbs.openhes.models.acquisition.ConnectionTypeSerialMoxa

Sub-message containing connection info for controlled-serial line (Moxa) connection type

## Fields

| Field | Information |
| --- | --- |
| host | <b>Type:</b> `string`<br><b>Description:</b><br>The host name or IP address of the device to connect to. |
| dataPort | <b>Type:</b> `uint32`<br><b>Description:</b><br>The TCP port number of the device to connect to - data port. |
| commandPort | <b>Type:</b> `uint32`<br><b>Description:</b><br>The TCP port number of the device to connect to - command port. |
| timeout | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The timeout for serial port connection that implements the Moxa protocol. |

