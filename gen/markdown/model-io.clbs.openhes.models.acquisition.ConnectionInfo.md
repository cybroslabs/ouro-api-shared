# Model: io.clbs.openhes.models.acquisition.ConnectionInfo

Sub-message containing connection info

## Fields

| Field | Type | Description |
| --- | --- | --- |
| tcpip | io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIp | The TCP connection type. |
| modemPool | io.clbs.openhes.models.acquisition.ConnectionTypeModemPool | The phone connection type. |
| serialOverIp | io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerial | The serial over IP connection type. |
| linkProtocol | io.clbs.openhes.models.acquisition.DataLinkProtocol | The data link protocol. |
| communicationBusId | string | The communication bus identifier. It behaves as a custom grouping key to link jobs together across multiple communication units. It shall be used in a situation when multiple entry points share single communication bus, e.g. multi-master RS-485 (primary and backup master). If not set then jobs are grouped by group-key defined based on the connection type. |
| attributes | map<string, io.clbs.openhes.models.common.FieldValue> | The connection attributes, see options in the DataLinkTemplate for given link_protocol. |

