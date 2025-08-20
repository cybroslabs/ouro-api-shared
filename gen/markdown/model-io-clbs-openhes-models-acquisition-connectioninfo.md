# Model: io.clbs.openhes.models.acquisition.ConnectionInfo

Defines the configuration parameters for device connection.

## Fields

| Field | Information |
| --- | --- |
| tcpip | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp`](model-io-clbs-openhes-models-acquisition-connectiontypedirecttcpip.md)<br><b>Description:</b><br>The TCP/IP connection type. |
| modemPool | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeModemPool`](model-io-clbs-openhes-models-acquisition-connectiontypemodempool.md)<br><b>Description:</b><br>The phone-based (modem pool) connection type.. |
| serialOverIp | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial`](model-io-clbs-openhes-models-acquisition-connectiontypecontrolledserial.md)<br><b>Description:</b><br>The serial-over-IP connection type. |
| linkProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DataLinkProtocol`](enum-io-clbs-openhes-models-acquisition-datalinkprotocol.md)<br><b>Description:</b><br>The data link protocol. |
| communicationBusId | <b>Type:</b> `string`<br><b>Description:</b><br>he communication bus identifier. Works as a custom grouping key to link jobs across multiple communication units when multiple entry points share a single communication bus (foe example, multi-master RS-485 withprimary and backup masters). If not set, jobs are grouped by the default group key based on the connection type. |
| attributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The connection attributes. See `GetDataLinkFields` in the acquisition package. |

