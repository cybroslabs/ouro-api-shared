# Model: io.clbs.openhes.models.acquisition.ActionResult

Defines the result of a single action.

## Fields

| Field | Information |
| --- | --- |
| actionId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique action identifier. |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionResultCode`](enum-io-clbs-openhes-models-acquisition-actionresultcode.md)<br><b>Description:</b><br>The status code of the action indicating success or failure. |
| data | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionData`](model-io-clbs-openhes-models-acquisition-actiondata.md)<br><b>Description:</b><br>The action result data containing the retrieved or processed information. |
| registerId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique register identifier. This is a read-only value and is set only if the action data relates to a register. Applicable only for results of regular bulks. |
| variableName | <b>Type:</b> `string`<br><b>Description:</b><br>The variable names. This is a read-only value and is set only if the action data relates to one or more variables based on device template mapping. Applicable only for results of regular bulks. |
| errorMessage | <b>Type:</b> [`io.clbs.openhes.models.common.FormattedMessage`](model-io-clbs-openhes-models-common-formattedmessage.md)<br><b>Description:</b><br>The user-facing error message provided when the action ends with `ERROR_CODE_ACTION_ERROR`. The message provides details about the failure, can be multiline, and should be in English. |

