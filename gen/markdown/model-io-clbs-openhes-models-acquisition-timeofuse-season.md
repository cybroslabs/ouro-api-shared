# Model: io.clbs.openhes.models.acquisition.timeofuse.Season

Season represents a season which spans across a specific start date and references a week

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>Unique identifier for the season<br><b>Example:</b> "season-winter" |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>Name of the season<br><b>Example:</b> "Winter" |
| startYear | <b>Type:</b> `int32`<br><b>Description:</b><br>Start year of the season<br><b>Example:</b> 2023 |
| startMonth | <b>Type:</b> `int32`<br><b>Description:</b><br>Start month of the season<br><b>Values:</b> 1–12<br><b>Example:</b> 12 |
| startDay | <b>Type:</b> `int32`<br><b>Description:</b><br>Start day of the season<br><b>Values:</b> 1–31<br><b>Example:</b> 1 |
| weekId | <b>Type:</b> `string`<br><b>Description:</b><br>Reference to a Week ID that this season is associated with<br><b>Example:</b> "week-01" |

