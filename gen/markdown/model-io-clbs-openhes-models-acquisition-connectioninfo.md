# Model: io.clbs.openhes.models.acquisition.ConnectionInfo

Sub-message containing connection info

## Fields

| Field | Information |
| --- | --- |
| tcpip | <b>Type:</b> [io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp](model-io-clbs-openhes-models-acquisition-connectiontypedirecttcpip.md)<br><b>Description:</b><br>The TCP connection type. |
| modemPool | <b>Type:</b> [io.clbs.openhes.models.acquisition.ConnectionTypeModemPool](model-io-clbs-openhes-models-acquisition-connectiontypemodempool.md)<br><b>Description:</b><br>The phone connection type. |
| serialOverIp | <b>Type:</b> [io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial](model-io-clbs-openhes-models-acquisition-connectiontypecontrolledserial.md)<br><b>Description:</b><br>The serial over IP connection type. |
| linkProtocol | <b>Type:</b> [io.clbs.openhes.models.acquisition.DataLinkProtocol](model-io-clbs-openhes-models-acquisition-datalinkprotocol.md)<br><b>Description:</b><br>The data link protocol. |
| communicationBusId | <b>Type:</b> string<br><b>Description:</b><br>The communication bus identifier. It behaves as a custom grouping key to link jobs together across multiple communication units. It shall be used in a situation when multiple entry points share single communication bus, e.g. multi-master RS-485 (primary and backup master). If not set then jobs are grouped by group-key defined based on the connection type. |
| attributes | <b>Type:</b> map<string, [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The connection attributes, see options in the DataLinkTemplate for given link_protocol. |

