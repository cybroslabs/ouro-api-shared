# Model: io.clbs.openhes.models.acquisition.timeofuse.DayProfile

DayProfile represents the profile for a single day, containing multiple Switching times

## Fields

| Field | Information |
| --- | --- |
| dayId | <b>Type:</b> `string`<br><b>Description:</b><br>Unique identifier for the day profile<br> <b>Example:</b> "weekday-profile" |
| switching | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.Switching`](model-io-clbs-openhes-models-acquisition-timeofuse-switching.md)<br><b>Description:</b><br>List of switching events (each with specific time and relay states) |

