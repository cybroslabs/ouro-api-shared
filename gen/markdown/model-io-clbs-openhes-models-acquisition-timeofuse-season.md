# Model: io.clbs.openhes.models.acquisition.timeofuse.Season

Defines a season that starts on a specific date and is linked to a week profile.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The unique season identifier.<br><b>Example:</b> "season-winter" |
| name | <b>Type:</b> `string`<br><b>Description:</b><br>The name of the season.<br><b>Example:</b> "Winter" |
| startYear | <b>Type:</b> `int32`<br><b>Description:</b><br>The start year of the season.<br><b>Example:</b> 2023 |
| startMonth | <b>Type:</b> `int32`<br><b>Description:</b><br>The start month of the season.<br><b>Values:</b> 1–12<br><b>Example:</b> 12 |
| startDay | <b>Type:</b> `int32`<br><b>Description:</b><br>The start day of the season.<br><b>Values:</b> 1–31<br><b>Example:</b> 1 |
| weekId | <b>Type:</b> `string`<br><b>Description:</b><br>The reference to a week ID that applies to this season.<br><b>Example:</b> "week-01" |

