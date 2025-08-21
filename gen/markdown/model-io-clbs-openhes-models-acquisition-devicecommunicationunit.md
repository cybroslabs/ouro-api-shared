# Model: io.clbs.openhes.models.acquisition.DeviceCommunicationUnit

Defines the mapping between a device and its communication unit.

## Fields

| Field | Information |
| --- | --- |
| communicationUnitId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the communication unit. |
| appProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](enum-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The application protocol used communication over the communication unit. |
| attributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The application protocol related device attributes, represented as a list of attribute definitions. |

