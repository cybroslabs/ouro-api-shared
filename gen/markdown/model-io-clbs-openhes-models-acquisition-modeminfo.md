# Model: io.clbs.openhes.models.acquisition.ModemInfo

Defines modem connection information.

## Fields

| Field | Information |
| --- | --- |
| modemId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the modem. It is automatically generated during creation. |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The name of the modem. |
| atInit | <b>Type:</b> `string`<br><b>Description:</b><br>The modem initialization command. For example, `AT&FE0X3`. |
| atDial | <b>Type:</b> `string`<br><b>Description:</b><br>The modem dial command. For example, `ATD`. |
| atHangup | <b>Type:</b> `string`<br><b>Description:</b><br>The modem hangup command. For example, `ATH`. |
| atEscape | <b>Type:</b> `string`<br><b>Description:</b><br>The modem escape command. For exampl, `+++`. |
| connectTimeout | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The modem connection timeout, if applicable by the `modem_connection` field. |
| commandTimeout | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The modem command timeout, if applicable given by the modem_connection field. |
| tcpip | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp`](model-io-clbs-openhes-models-acquisition-connectiontypedirecttcpip.md)<br><b>Description:</b><br>The TCP connection type. The modem either supports TCP directly, or an IP-to-serial converter is used that handles the serial configuration. |
| serialOverIp | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial`](model-io-clbs-openhes-models-acquisition-connectiontypecontrolledserial.md)<br><b>Description:</b><br>The serial-over-IP connection type. Used when the modem is connected via an IP-to-serial converter and requires specific handling. |
| serialBaudRate | <b>Type:</b> `uint32`<br><b>Description:</b><br>The serial baud rate, if applicable by the `modem_connection` field. |

