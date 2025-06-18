# Enum: io.clbs.openhes.models.common.FieldDisplayFormat

Enum representing the field display format.

## Options

| Value | Description |
| --- | --- |
| DEFAULT | The default display format. The GUI shows text, int, double, date-time or duration in the default format. |
| DATE_ONLY | The date-only display format of local date-time. Data type must be TIMESTAMP or string. |
| UTC_DATETIME | The date-time display format with in UTC timezone. Data type must be TIMESTAMP or string. |
| UTC_DATE_ONLY | The date display format with in UTC timezone. Data type must be TIMESTAMP or string. |
| TIMEOFDAY | The time of day display format hh:mm:ss, e.g. 12:30:00. Data type must be INTEGER in milliseconds. |
| MONEY | The money display format. The unit must be set to the currency code, ISO 4217 standard (e.g. USD, EUR, ...). Data type must be DOUBLE or INTEGER. |
| PASSWORD | The password display format. Data type must be TEXT. The GUI must always display six starts (******) not to reveal the actual password length. |
| MULTILINE | The multiline-string display format. Data type must be TEXT. |
