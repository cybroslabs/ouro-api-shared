# Model: io.clbs.openhes.models.acquisition.DataLinkTemplate

Sub-message containing destription for one data link protocol, e.g. HDLC.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| linkProtocol | [io.clbs.openhes.models.acquisition.DataLinkProtocol](model-io-clbs-openhes-models-acquisition-datalinkprotocol.md) | The data link protocol. |
| appProtocolRefs | [io.clbs.openhes.models.acquisition.ApplicationProtocol](model-io-clbs-openhes-models-acquisition-applicationprotocol.md) | The list of application protocol identifiers supported by the driver. |
| attributes | [io.clbs.openhes.models.common.FieldDescriptor](model-io-clbs-openhes-models-common-fielddescriptor.md) | The list of attribute definitions related to given data link type (see link_protocol property) and all those will be instantiated for each device. |

