# Model: io.clbs.openhes.models.acquisition.ActionResult

Sub-message containing action result for a single action.

## Fields

| Field | Information |
| --- | --- |
| actionId | <b>Type:</b> `string`<br><b>Description:</b><br>The action identifier. |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionResultCode`](enum-io-clbs-openhes-models-acquisition-actionresultcode.md)<br><b>Description:</b><br>The status of the action. |
| data | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionData`](model-io-clbs-openhes-models-acquisition-actiondata.md)<br><b>Description:</b><br>The action result data. |
| registerId | <b>Type:</b> `string`<br><b>Description:</b><br>The register identifier. It is a read-only value, set only if the action data are related to a register. The value is applicable only and only for results for regular bulks. |
| variableName | <b>Type:</b> `string`<br><b>Description:</b><br>The variable name(s). It is a read-only value, set only if the action data are related to one or more variables based on device template mapping. The value is applicable only and only for results for regular bulks. |
| errorMessage | <b>Type:</b> [`io.clbs.openhes.models.common.FormattedMessage`](model-io-clbs-openhes-models-common-formattedmessage.md)<br><b>Description:</b><br>The user error message related to the action which ends up with code ERROR_CODE_ACTION_ERROR, if any. It is used to provide more details about the action result. It can be multiline and should be in English language. |

