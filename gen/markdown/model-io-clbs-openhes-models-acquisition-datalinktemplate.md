# Model: io.clbs.openhes.models.acquisition.DataLinkTemplate

Sub-message containing destription for one data link protocol, e.g. HDLC.

## Fields

| Field | Information |
| --- | --- |
| linkProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DataLinkProtocol`](enum-io-clbs-openhes-models-acquisition-datalinkprotocol.md)<br><b>Description:</b><br>The data link protocol. |
| appProtocolRefs | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](enum-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The list of application protocol identifiers supported by the driver. |
| attributes | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)<br><b>Description:</b><br>The list of attribute definitions related to given data link type (see link_protocol property). The field definitions are taken from the system, drivers must leave this empty. |

