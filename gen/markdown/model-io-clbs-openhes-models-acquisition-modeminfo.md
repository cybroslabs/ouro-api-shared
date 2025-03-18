# Model: io.clbs.openhes.models.acquisition.ModemInfo

Sub-message containing modem connection info

## Fields

| Field | Information |
| --- | --- |
| modemId | <b>Type:</b> string<br><b>Description:</b><br>The modem identifier. It is automatically generated during creation. |
| name | <b>Type:</b> string<br><b>Description:</b><br>The name of the modem. |
| atInit | <b>Type:</b> string<br><b>Description:</b><br>The modem initialization command, e.g. AT&FE0X3 |
| atDial | <b>Type:</b> string<br><b>Description:</b><br>The modem dial command, e.g. ATD. |
| atHangup | <b>Type:</b> string<br><b>Description:</b><br>The modem hangup command, e.g. ATH. |
| atEscape | <b>Type:</b> string<br><b>Description:</b><br>The modem escape command, e.g. +++. |
| connectTimeout | <b>Type:</b> google.protobuf.Duration<br><b>Description:</b><br>The modem connection timeout, if applicable given by the modem_connection field. |
| commandTimeout | <b>Type:</b> google.protobuf.Duration<br><b>Description:</b><br>The modem command timeout, if applicable given by the modem_connection field. |
| tcpip | <b>Type:</b> [io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp](model-io-clbs-openhes-models-acquisition-connectiontypedirecttcpip.md)<br><b>Description:</b><br>The TCP connection type. The modem has either TCP or there is a IP-to-serial converter which handles the serial configuration so no additional serial configuration is needed. |
| serialOverIp | <b>Type:</b> [io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial](model-io-clbs-openhes-models-acquisition-connectiontypecontrolledserial.md)<br><b>Description:</b><br>The serial over IP connection type. The modem is connected behind an IP-to-serial converter and needs connection specific handling. |
| serialBaudRate | <b>Type:</b> uint32<br><b>Description:</b><br>The serial baud rate, if applicable given by the modem_connection field. |

