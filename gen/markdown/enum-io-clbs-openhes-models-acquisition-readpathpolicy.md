# Enum: io.clbs.openhes.models.acquisition.ReadPathPolicy

Defines the read path policies for retrieving data from devices.

## Options

| Value | Description |
| --- | --- |
| READ_PATH_POLICY_UNSPECIFIED | Keeps the current behavior (typically data-concentrator first). |
| READ_PATH_POLICY_METER_FIRST | Attempts to read directly from the meter and falls back to the data concentrator if needed and supported. |
| READ_PATH_POLICY_DC_FIRST | Attempts to read from the data concentrator first and falls back to the meter if needed and supported. |
