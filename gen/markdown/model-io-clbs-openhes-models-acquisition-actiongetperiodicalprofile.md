# Model: io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfile

Sub-message containing get periodical profile action specification

## Fields

| Field | Information |
| --- | --- |
| from | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The start timestamp of the profile. |
| to | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The end timestamp of the profile. |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the target register, e.g. integer, double, string, timestamp, etc. This value can be specified for proxy bulk, while it's automatically taken from the register definition in the system for bulks. |

