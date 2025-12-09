# Model: io.clbs.openhes.models.acquisition.ModemInfo

Defines modem connection information.

## Fields

| Field | Information |
| --- | --- |
| modemId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique identifier of the modem. It is automatically generated during creation. |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The name of the modem.<br><b>Example:</b> "GSM-Gateway-Main" |
| atInit | <b>Type:</b> `string`<br><b>Description:</b><br>The modem initialization command.<br><b>Example:</b> "ATZ" |
| atDial | <b>Type:</b> `string`<br><b>Description:</b><br>The modem dial command.<br><b>Example:</b> "ATDT" |
| atHangup | <b>Type:</b> `string`<br><b>Description:</b><br>The modem hangup command.<br><b>Example:</b> "ATH0" |
| atEscape | <b>Type:</b> `string`<br><b>Description:</b><br>The modem escape command.<br><b>Example:</b> "+++" |
| connectTimeout | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The modem connection timeout, if applicable by the `modem_connection` field. |
| commandTimeout | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The modem command timeout, if applicable given by the modem_connection field. |
| tcpip | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp`](model-io-clbs-openhes-models-acquisition-connectiontypedirecttcpip.md)<br><b>Description:</b><br>The TCP connection type. The modem either supports TCP directly, or an IP-to-serial converter is used that handles the serial configuration. |
| serialOverIp | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial`](model-io-clbs-openhes-models-acquisition-connectiontypecontrolledserial.md)<br><b>Description:</b><br>The serial-over-IP connection type. Used when the modem is connected via an IP-to-serial converter and requires specific handling. |

