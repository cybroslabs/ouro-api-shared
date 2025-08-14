# Enum: io.clbs.openhes.models.acquisition.ReadPathPolicy

## Options

| Value | Description |
| --- | --- |
| READ_PATH_POLICY_UNSPECIFIED | Keep current behavior (likely data-concentrator first). |
| METER_FIRST | Try meter directly; fall back to data-concentrator if needed and supported. |
| DC_FIRST | Try data-concentrator first; fall back to meter if needed and supported. |
