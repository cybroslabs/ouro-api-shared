# Model: io.clbs.openhes.models.acquisition.JobDevice

Defines single job-device job-device information.

## Fields

| Field | Information |
| --- | --- |
| jobId | <b>Type:</b> `string`<br><b>Description:</b><br>The device job identifier within the parent bulk. |
| deviceId | <b>Type:</b> `string`<br><b>Description:</b><br>The device identifier. If set, all subsequent details are loaded from the device registry. |
| externalId | <b>Type:</b> `string`<br><b>Description:</b><br>The external identifier of the device. |
| deviceAttributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The connection attributes for the device. Options depend on the `ApplicationProtocolTemplate` of the given application protocol (see the `app_protocol` property). |
| connectionInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionInfo`](model-io-clbs-openhes-models-acquisition-connectioninfo.md)<br><b>Description:</b><br>The device connection parameters. |
| appProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](enum-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The application protocol used by the device. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone associated with the device. For example, `"America/New_York"`, `"Europe/Prague"`, `"CET"`, `"GMT"`, `"Etc/GMT+2"`. |

