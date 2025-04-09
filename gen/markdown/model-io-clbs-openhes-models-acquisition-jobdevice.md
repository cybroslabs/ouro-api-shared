# Model: io.clbs.openhes.models.acquisition.JobDevice

Sub-message representing a single job-device info.

## Fields

| Field | Information |
| --- | --- |
| jobId | <b>Type:</b> `string`<br><b>Description:</b><br>The device's job identifier within the parent bulk. |
| deviceId | <b>Type:</b> `string`<br><b>Description:</b><br>The device identifier. If set then all below is loaded from the device registry. |
| externalId | <b>Type:</b> `string`<br><b>Description:</b><br>The external identifier. |
| deviceAttributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The connection attributes to the device, see options in the ApplicationProtocolTemplate for given application protocol (see app_protocol property). |
| connectionInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionInfo`](model-io-clbs-openhes-models-acquisition-connectioninfo.md)<br><b>Description:</b><br>The connection (device) parameters. |
| appProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](model-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The application protocol. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone related to the device, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2". |

