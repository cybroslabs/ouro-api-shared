# Model: io.clbs.openhes.models.acquisition.ConnectionTypeSerialMoxa

Sub-message containing connection info for controlled-serial line (Moxa) connection type

## Fields

| Field | Type | Description |
| --- | --- | --- |
| host | string | The host name or IP address of the device to connect to. |
| dataPort | uint32 | The TCP port number of the device to connect to - data port. |
| commandPort | uint32 | The TCP port number of the device to connect to - command port. |
| timeout | google.protobuf.Duration | The timeout for serial port connection that implements the Moxa protocol. |

