# Model: io.clbs.openhes.models.acquisition.timeofuse.Switching

Switching data includes tariffs and relay states for specific times

## Fields

| Field | Information |
| --- | --- |
| hour | <b>Type:</b> `int32`<br><b>Description:</b><br>Hour of the switching event.<br><b>Values:</b> 0–23<br><b>Example:</b> 7 |
| minute | <b>Type:</b> `int32`<br><b>Description:</b><br>Minute of the switching event.<br><b>Values:</b> 0–59<br><b>Example:</b> 30 |
| tariff | <b>Type:</b> `int32`<br><b>Description:</b><br>Tariff ID.<br><b>Values:</b> -1 means no tariff; other values are valid tariff IDs<br><b>Example:</b> 2 |
| relays | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.RelayStateRecord`](model-io-clbs-openhes-models-acquisition-timeofuse-relaystaterecord.md)<br><b>Description:</b><br>Map of relay ID to relay state. The list must not contain duplicit relay IDs. |

