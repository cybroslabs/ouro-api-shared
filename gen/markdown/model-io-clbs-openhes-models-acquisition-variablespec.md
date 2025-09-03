# Model: io.clbs.openhes.models.acquisition.VariableSpec

Defines a variable specification.

## Fields

| Field | Information |
| --- | --- |
| registerId | <b>Type:</b> `string`<br><b>Description:</b><br>The list of device configuration register identifiers. |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the variable. Only registers of the same data type can be linked to the variable. |
| excludeDataFrom | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether the returned data for this variabe should exclude values at `from` timestamp. If set to `true`, the returned device data returned for this variable dos not include values at `from` timestamp. Default value is `false`. |

