# Enum: io.clbs.openhes.models.acquisition.ReadPathPolicy

Defines the read path policies for retrieving data from devices.

## Options

| Value | Description |
| --- | --- |
| READ_PATH_POLICY_UNSPECIFIED | Keeps the current behavior (typically data-concentrator first). |
| METER_FIRST | Attempts to read directly from the meter and falls back to the data concentrator if needed and supported. |
| DC_FIRST | Attempts to read from the data concentrator first and falls back to the meter if needed and supported. |
