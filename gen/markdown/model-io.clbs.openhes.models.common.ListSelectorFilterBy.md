# Model: io.clbs.openhes.models.common.ListSelectorFilterBy

The filtering criteria.

 Depending on the operator, the 'text', 'integer', 'number', 'boolean' or 'date' field should be used.
 - No value must be set for operators: 'IS_NULL', 'IS_NOT_NULL'.
 - One value must be set for single operand operators: 'EQUAL', 'NOT_EQUAL', 'GREATER_THAN', 'GREATER_THAN_OR_EQUAL', 'LESS_THAN', 'LESS_THAN_OR_EQUAL', 'CONTAINS', 'NOT_CONTAINS', 'STARTS_WITH', 'ENDS_WITH'.
 - Two values must be set for two operand operators: 'BETWEEN'.
 - Any number of values can be set for generic operators: 'IN', 'NOT_IN'.

 Field type determines the data type and only related field should be used. Other fields shall not be set and will be ignored by the system.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| fieldId | string | Field id. |
| operator | io.clbs.openhes.models.common.FilterOperator | The filter operator. |
| dataType | io.clbs.openhes.models.common.FieldDataType | The data type of the field. |
| text | string | The text-typed value(s) used for filtering. |
| integer | sint64 | The integer-typed value(s) used for filtering. |
| number | double | The number-typed value(s) used for filtering. |
| boolean | bool | The boolean-typed value(s) used for filtering. |
| date | google.protobuf.Timestamp | The date-typed value(s) used for filtering. |

