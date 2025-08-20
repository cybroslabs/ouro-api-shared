# Model: io.clbs.openhes.models.acquisition.DataLinkTemplate

Defines the destription of a single data link protocol, for example `HDLC`.

## Fields

| Field | Information |
| --- | --- |
| linkProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DataLinkProtocol`](enum-io-clbs-openhes-models-acquisition-datalinkprotocol.md)<br><b>Description:</b><br>The data link protocol. |
| appProtocolRefs | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](enum-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The list of application protocol identifiers supported by the driver. |
| attributes | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)<br><b>Description:</b><br>The list of attribute definitions related to the selected data link type (see l`ink_protocol` property). These field definitions are provided by the system and drivers must leave this field empty. |

