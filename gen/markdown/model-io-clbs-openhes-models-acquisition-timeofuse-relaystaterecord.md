# Model: io.clbs.openhes.models.acquisition.timeofuse.RelayStateRecord

RelayStateRecord represents the state of a relay at a specific time.
 It contains the relay ID and its state (CONNECT or DISCONNECT).
 The relay ID must be unique within the list of relays.

## Fields

| Field | Information |
| --- | --- |
| relayId | <b>Type:</b> `int32`<br><b>Description:</b><br>Relay ID<br> <b>Example:</b> 1 |
| state | <b>Type:</b> [`io.clbs.openhes.models.acquisition.timeofuse.RelayState`](model-io-clbs-openhes-models-acquisition-timeofuse-relaystate.md)<br><b>Description:</b><br>State of the relay (CONNECT or DISCONNECT)<br> <b>Example:</b> CONNECT |

