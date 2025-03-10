# Model: io.clbs.openhes.models.acquisition.JobDevice

Sub-message representing a single job-device info.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| jobId | string | The device's job identifier within the parent bulk. |
| deviceId | string | The device identifier. If set then all below is loaded from the device registry. |
| externalId | string | The external identifier. |
| deviceAttributes | map<string, [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md)> | The connection attributes to the device. |
| connectionInfo | [io.clbs.openhes.models.acquisition.ConnectionInfo](model-io-clbs-openhes-models-acquisition-connectioninfo.md) | The connection (device) parameters. |
| appProtocol | [io.clbs.openhes.models.acquisition.ApplicationProtocol](model-io-clbs-openhes-models-acquisition-applicationprotocol.md) | The application protocol. |
| timezone | string | The timezone related to the device, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2". |

