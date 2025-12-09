# Model: io.clbs.openhes.models.acquisition.ActionGetIrregularProfile

Defines the get irregular profile action specification.

## Fields

| Field | Information |
| --- | --- |
| rangeStart | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the profile readout (inclusive). |
| rangeEnd | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the profile readout (inclusive). |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the target register, for example `INTEGER`, `DOUBLE`, `TEXT`, `TIMESTAMP`. For proxy bulks, this value can be explicitly specified. For regular bulks, it is automatically taken from the register definition in the system. |

