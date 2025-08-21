# Model: io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfile

Defines the get periodical profile action specification.

## Fields

| Field | Information |
| --- | --- |
| rangeStart | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the profile readout. |
| rangeEnd | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the profile readout. |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the target register, for example `integer`, `double`, `string`, `timestamp`. For proxy bulks, this value can be explicitely specified. For regular bulks, it is automatically taken from the register definition in the system. |

