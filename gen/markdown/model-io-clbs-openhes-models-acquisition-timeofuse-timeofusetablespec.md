# Model: io.clbs.openhes.models.acquisition.timeofuse.TimeOfUseTableSpec

TimeOfUse represents the main Time-of-Use (TOU) table containing all relevant definitions

## Fields

| Field | Information |
| --- | --- |
| expiesAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>Expiration date of the TOU table<br> <b>Example:</b> "2025-01-01T00:00:00Z" |
| hdoGroupId | <b>Type:</b> `string`<br><b>Description:</b><br>HDO (High Demand Option) Group ID<br> <b>Example:</b> "group-a" |
| activateAt | <b>Type:</b> `google.type.Date`<br><b>Description:</b><br>Activation date of the TOU table<br> <b>Example:</b> "2024-06-01" |
| seasons | <b>Type:</b> `[io.clbs.openhes.models.acquisition.timeofuse.Season](model-io-clbs-openhes-models-acquisition-timeofuse-season.md)`<br><b>Description:</b><br>List of seasons defined in the TOU |
| weeks | <b>Type:</b> `[io.clbs.openhes.models.acquisition.timeofuse.Week](model-io-clbs-openhes-models-acquisition-timeofuse-week.md)`<br><b>Description:</b><br>List of weeks defined in the TOU |
| dayProfiles | <b>Type:</b> `[io.clbs.openhes.models.acquisition.timeofuse.DayProfile](model-io-clbs-openhes-models-acquisition-timeofuse-dayprofile.md)`<br><b>Description:</b><br>List of day profiles (each day having a list of switching events) |
| specialDays | <b>Type:</b> `[io.clbs.openhes.models.acquisition.timeofuse.SpecialDay](model-io-clbs-openhes-models-acquisition-timeofuse-specialday.md)`<br><b>Description:</b><br>List of special days (e.g., holidays, exceptions) |

