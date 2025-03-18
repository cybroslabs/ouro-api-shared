# Model: io.clbs.openhes.models.acquisition.ConnectionTypeSerialRfc2217

Sub-message containing connection info for controlled-serial line (RFC 2217) connection type

## Fields

| Field | Information |
| --- | --- |
| host | <b>Type:</b> string<br><b>Description:</b><br>The host name or IP address of the device to connect to. |
| port | <b>Type:</b> uint32<br><b>Description:</b><br>The TCP port number of the device to connect to. |
| timeout | <b>Type:</b> google.protobuf.Duration<br><b>Description:</b><br>The timeout for serial port connection that implements the RFC 2217 protocol. |

