# Model: io.clbs.openhes.models.acquisition.JobDevice

Defines the information for a device job.

## Fields

| Field | Information |
| --- | --- |
| jobId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique job identifier within the parent bulk. |
| deviceId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique device identifier. If set, all subsequent details are loaded automatically from the device registry. |
| externalId | <b>Type:</b> `string`<br><b>Description:</b><br>The external device identifier.<br><b>Example:</b> "EXT-001" |
| deviceAttributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The connection attributes for the device. The options depend on the `ApplicationProtocolTemplate` of the given application protocol (see the `app_protocol` property). |
| connectionInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ConnectionInfo`](model-io-clbs-openhes-models-acquisition-connectioninfo.md)<br><b>Description:</b><br>The device connection parameters. |
| appProtocol | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ApplicationProtocol`](enum-io-clbs-openhes-models-acquisition-applicationprotocol.md)<br><b>Description:</b><br>The application protocol used to communicate with the device. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone associated with the device.<br><b>Example:</b> "Etc/GMT+2" |

