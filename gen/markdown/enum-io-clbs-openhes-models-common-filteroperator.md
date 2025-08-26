# Enum: io.clbs.openhes.models.common.FilterOperator

Defines the operators available for filtering options.

## Options

| Value | Description |
| --- | --- |
| FILTER_OPERATOR_UNSPECIFIED | Unspecified filter operator. |
| EQUAL | Single-operand operator. Matches fields equal to the provided value. Supported for: `text`, `integer`, `number`, `boolean`, `date` fields. |
| NOT_EQUAL | Single-operand operator. Matches fields not equal to the provided value. Supported for: `text`, `integer`, `number`, `boolean`, `date` fields. |
| GREATER_THAN | Single-operand operator. Matches fields greater than the provided value. Supported for: `integer`, `number`, `date` fields. |
| GREATER_THAN_OR_EQUAL | Single-operand operator. Matches fields greater than or equal to the provided value. Supported for: `integer`, `number`, `date` fields. |
| LESS_THAN | Single-operand operator. Matches fields less than the provided value. Supported for: `integer`, `number`, `date` fields. |
| LESS_THAN_OR_EQUAL | Single-operand operator. Matches fields less than or equal to the provided value. Supported for: `integer`, `number`, `date` fields. |
| CONTAINS | Single-operand operator. Matches text fields that contain the provided string. Supported for: `text` fields. |
| NOT_CONTAINS | Single-operand operator. Matches text fields that do not contain the provided string. Supported for: `text` fields. |
| STARTS_WITH | Single-operand operator. Matches text fields that start with the provided string. Supported for: `text` fields. |
| ENDS_WITH | Single-operand operator. Matches text fields that end with the provided substring. Supported for: `text` fields. |
| IN | Multi-operand operator. Matches fields whose value is within the provided set. Supported for: `text`, `integer`, `number`, `boolean` fields. |
| NOT_IN | Multi-operand operator. Matches fields whose value is not within the provided set. Supported for: `text`, `integer`, `number`, `boolean` fields. |
| BETWEEN | Two-operand operator. Matches fields with values between the two provided operands (inclusive). Supported for: `integer`, `number`, `date` fields. |
| IS_NULL | No-operand operator. Matches fields that are either `null` or an empty string. |
| IS_NOT_NULL | No-operand operator. Matches fields that are not `null` or an empty string. |
