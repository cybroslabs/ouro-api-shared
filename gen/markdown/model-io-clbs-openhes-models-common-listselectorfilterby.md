# Model: io.clbs.openhes.models.common.ListSelectorFilterBy

Defines the filtering criteria for listing objects.

 Depending on the operator, a specific filed type must be provided.
 - **No value** is required for operators: `IS_NULL`, `IS_NOT_NULL`.
 - Exactly **one value** must be set for single-operand operators: `EQUAL`, `NOT_EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `CONTAINS`, `NOT_CONTAINS`, `STARTS_WITH`, `ENDS_WITH`.
 - Exactly **two values** must be set for two-operand operators: `BETWEEN`.
 - **Zero or more** values can be set for generic operators: `IN`, `NOT_IN`.

 Only the field corresponding to th declared data type must be set. Other fields must not be set will be ignored by the system.

## Fields

| Field | Information |
| --- | --- |
| path | <b>Type:</b> `string`<br><b>Description:</b><br>The field path to filter by. This may be a JSON path (`js_path`) or a path from `FieldDescriptor`. |
| operator | <b>Type:</b> [`io.clbs.openhes.models.common.FilterOperator`](enum-io-clbs-openhes-models-common-filteroperator.md)<br><b>Description:</b><br>The operator that defines the filtering condition. |
| dataType | <b>Type:</b> [`io.clbs.openhes.models.common.FieldDataType`](enum-io-clbs-openhes-models-common-fielddatatype.md)<br><b>Description:</b><br>The data type of the field. |
| text | <b>Type:</b> `string`<br><b>Description:</b><br>The values for string-based filtering. |
| integer | <b>Type:</b> `sint64`<br><b>Description:</b><br>The values for integer-based filtering. |
| number | <b>Type:</b> `double`<br><b>Description:</b><br>The values for numeric filtering. |
| boolean | <b>Type:</b> `bool`<br><b>Description:</b><br>The values for boolean-based filtering. |
| date | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>Values for date-based filtering. |

