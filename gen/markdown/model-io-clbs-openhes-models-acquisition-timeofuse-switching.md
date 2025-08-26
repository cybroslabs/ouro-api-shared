# Model: io.clbs.openhes.models.acquisition.timeofuse.Switching

Defines the switching configuration, including tariffs and relay states for specific times.

## Fields

| Field | Information |
| --- | --- |
| hour | <b>Type:</b> `int32`<br><b>Description:</b><br>The hour of the switching event.<br><b>Values:</b> 0–23<br><b>Example:</b> 7 |
| minute | <b>Type:</b> `int32`<br><b>Description:</b><br>The minute of the switching event.<br><b>Values:</b> 0–59<br><b>Example:</b> 30 |
| tariff | <b>Type:</b> `int32`<br><b>Description:</b><br>The tariff ID.<br><b>Values:</b> -1 means no tariff; other values are valid tariff IDs<br><b>Example:</b> 2 |
| relays | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.RelayStateRecord`](model-io-clbs-openhes-models-acquisition-timeofuse-relaystaterecord.md)<br><b>Description:</b><br>A map of relay IDs to their corresponding relay states. Duplicate IDs are not allowed. |

