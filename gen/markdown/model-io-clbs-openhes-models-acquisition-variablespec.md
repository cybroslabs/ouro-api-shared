# Model: io.clbs.openhes.models.acquisition.VariableSpec

## Fields

| Field | Information |
| --- | --- |
| registerId | <b>Type:</b> `string`<br><b>Description:</b><br>The list of device configuration register identifiers. |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the variable. Only registers of the same data type can be linked to the variable. |
| excludeDataFrom | <b>Type:</b> `bool`<br><b>Description:</b><br>If true, the device data returned for this variable will exclude values at 'from' timestamp. Default is false (include values at 'from' timestamp). |

