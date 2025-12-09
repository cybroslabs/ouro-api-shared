# Model: io.clbs.openhes.models.acquisition.DataLinkTemplate

Defines the template description of a single data link protocol, for example HDLC.

## Fields

| Field | Information |
| --- | --- |
| linkProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DataLinkProtocol`](enum-io-clbs-openhes-models-acquisition-datalinkprotocol.md)<br><b>Description:</b><br>The data link protocol type. |
| appProtocolRefs | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](enum-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The list of application protocol identifiers supported by the driver for this data link. |
| attributes | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)<br><b>Description:</b><br>The list of attribute definitions related to the selected data link type (see `link_protocol` property). These field definitions are provided by the system and drivers must leave this field empty. |

