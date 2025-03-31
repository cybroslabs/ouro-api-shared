# Model: io.clbs.openhes.models.acquisition.ActionResult

Sub-message containing action result for a single action.

## Fields

| Field | Information |
| --- | --- |
| actionId | <b>Type:</b> string<br><b>Description:</b><br>The action identifier. |
| status | <b>Type:</b> [io.clbs.openhes.models.acquisition.ActionResultCode](model-io-clbs-openhes-models-acquisition-actionresultcode.md)<br><b>Description:</b><br>The status of the action. |
| data | <b>Type:</b> [io.clbs.openhes.models.acquisition.ActionData](model-io-clbs-openhes-models-acquisition-actiondata.md)<br><b>Description:</b><br>The action result data. |
| registerId | <b>Type:</b> string<br><b>Description:</b><br>The register identifier. It is a read-only value, set only if the action data are related to a register. The value is applicable only and only for regular bulks. |
| variableId | <b>Type:</b> string<br><b>Description:</b><br>The variable identifier. It is a read-only value, set only if the action data are related to a variable. The value is applicable only and only for regular bulks. |

