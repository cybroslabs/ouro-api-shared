# Enum: io.clbs.openhes.models.common.FieldDisplayFormat

Define the display format, determining how alues are presetned in the GUI.

## Options

| Value | Description |
| --- | --- |
| DISPLAY_FORMAT_UNSPECIFIED | The default display format. Text, integer, double, date-time, or duration values are shown in their default format. |
| DATE_ONLY | Displays only the date portion of a local date-time value. Data type must be `TIMESTAMP` or `string`. |
| UTC_DATETIME | Displays the full date-time in UTC timezone. Data type must be `TIMESTAMP` or `string`. |
| UTC_DATE_ONLY | Displays the date in UTC timezone. Data type must be `TIMESTAMP` or `string`. |
| TIMEOFDAY | Displays time of day in `hh:mm:ss` format (for example, `12:30:00`). Data type must be `INTEGER` in milliseconds. |
| MONEY | Displays money amounts. The unit must be set to an ISO 4217 currency code (for example, USD, EUR). Data type must be `DOUBLE` or `INTEGER`. |
| PASSWORD | Displays values as hidden passwords. Data type must be `TEXT`. The GUI always shows six asterisks (******) regardless of actual length. |
| MULTILINE | Displays values as multi-line text. Data type must be `TEXT`. |
| COMBO | Displays values as a combo-box with selectable options. Data type must be `TEXT`. |
