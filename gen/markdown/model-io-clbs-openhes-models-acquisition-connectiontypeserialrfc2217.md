# Model: io.clbs.openhes.models.acquisition.ConnectionTypeSerialRfc2217

Sub-message containing connection info for controlled-serial line (RFC 2217) connection type

## Fields

| Field | Type | Description |
| --- | --- | --- |
| host | string | The host name or IP address of the device to connect to. |
| port | uint32 | The TCP port number of the device to connect to. |
| timeout | google.protobuf.Duration | The timeout for serial port connection that implements the RFC 2217 protocol. |

