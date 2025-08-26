# Model: io.clbs.openhes.models.acquisition.timeofuse.DayProfile

Defines a single day profile, containing multiple switching times.

## Fields

| Field | Information |
| --- | --- |
| dayId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique day profile identifier.<br><b>Example:</b> "weekday-profile" |
| switching | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.Switching`](model-io-clbs-openhes-models-acquisition-timeofuse-switching.md)<br><b>Description:</b><br>The list of switching events. Each event specifies the time and relay states. |

