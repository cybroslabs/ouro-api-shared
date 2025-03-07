# Model: io.clbs.openhes.models.acquisition.ModemInfo

Sub-message containing modem connection info

## Fields

| Field | Type | Description |
| --- | --- | --- |
| modemId | string | The modem identifier. It is automatically generated during creation. |
| name | string | The name of the modem. |
| atInit | string | The modem initialization command, e.g. AT&FE0X3 |
| atDial | string | The modem dial command, e.g. ATD. |
| atHangup | string | The modem hangup command, e.g. ATH. |
| atEscape | string | The modem escape command, e.g. +++. |
| connectTimeout | google.protobuf.Duration | The modem connection timeout, if applicable given by the modem_connection field. |
| commandTimeout | google.protobuf.Duration | The modem command timeout, if applicable given by the modem_connection field. |
| tcpip | io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp | The TCP connection type. The modem has either TCP or there is a IP-to-serial converter which handles the serial configuration so no additional serial configuration is needed. |
| serialOverIp | io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial | The serial over IP connection type. The modem is connected behind an IP-to-serial converter and needs connection specific handling. |
| serialBaudRate | uint32 | The serial baud rate, if applicable given by the modem_connection field. |

