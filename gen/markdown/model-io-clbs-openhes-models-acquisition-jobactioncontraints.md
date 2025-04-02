# Model: io.clbs.openhes.models.acquisition.JobActionContraints

Sub-message containing constraints for job actions.

## Fields

| Field | Information |
| --- | --- |
| getRegisterTypeName | <b>Type:</b> `map<string, string>`<br><b>Description:</b><br>Contains all the register types & their names in the English language. For example: [gen]="Generic Register", [vqi]="Quality VQI Register"<br> Action: ACTION_TYPE_GET_REGISTER |
| getRegisterTypeAttributes | <b>Type:</b> `map<string, [io.clbs.openhes.models.common.ListOfString](model-io-clbs-openhes-models-common-listofstring.md)>`<br><b>Description:</b><br>Contains all the register types & their attributes. For example: [gen]=["attribute1", "attribute2"], [vqi]=["attribute1", "attribute3"]<br> See JobActionAttributes for the attribute definitions.<br> Action: ACTION_TYPE_GET_REGISTER |

