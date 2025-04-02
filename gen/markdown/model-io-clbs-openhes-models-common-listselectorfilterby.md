# Model: io.clbs.openhes.models.common.ListSelectorFilterBy

The filtering criteria.

 Depending on the operator, the `text`, `integer`, `number`, `boolean` or `date` field should be used.
 - **No value** must be set for operators: `IS_NULL`, `IS_NOT_NULL`.
 - Exactly **One value** must be set for single operand operators: `EQUAL`, `NOT_EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `CONTAINS`, `NOT_CONTAINS`, `STARTS_WITH`, `ENDS_WITH`.
 - Exactly **Two values** must be set for two operand operators: `BETWEEN`.
 - **Zero or more** values can be set for generic operators: `IN`, `NOT_IN`.

 Field type determines the data type and only related field should be used. Other fields shall not be set and will be ignored by the system.

## Fields

| Field | Information |
| --- | --- |
| fieldId | <b>Type:</b> string<br><b>Description:</b><br>Field id. |
| operator | <b>Type:</b> [io.clbs.openhes.models.common.FilterOperator](model-io-clbs-openhes-models-common-filteroperator.md)<br><b>Description:</b><br>The filter operator. |
| dataType | <b>Type:</b> [io.clbs.openhes.models.common.FieldDataType](model-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the field. |
| text | <b>Type:</b> string<br><b>Description:</b><br>The text-typed value(s) used for filtering. |
| integer | <b>Type:</b> sint64<br><b>Description:</b><br>The integer-typed value(s) used for filtering. |
| number | <b>Type:</b> double<br><b>Description:</b><br>The number-typed value(s) used for filtering. |
| boolean | <b>Type:</b> bool<br><b>Description:</b><br>The boolean-typed value(s) used for filtering. |
| date | <b>Type:</b> google.protobuf.Timestamp<br><b>Description:</b><br>The date-typed value(s) used for filtering. |

