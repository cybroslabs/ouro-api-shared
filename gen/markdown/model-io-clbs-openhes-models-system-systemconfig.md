# Model: io.clbs.openhes.models.system.SystemConfig

## Fields

| Field | Information |
| --- | --- |
| maxReplicas | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of replicas for all drivers.<br>    0 represents no active replicas will run, effectively disabling acquisition<br>   >0 represents the maximum number of replicas per driver |
| maxCascadeDeviceCount | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of cascade devices for the driver. Minimum is 1. |
| maxSlotsPerDriver | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of slots per driver<br>   -1 represents unlimited number of slots, effecticaly using maximum number of slots supported by driver<br>    0 represents no active slots will run, effectively disabling acquisition<br>   >0 represents the maximum number of slots per driver, the number of slots never exceeds the number of slots supported by driver |
| minReplicas | <b>Type:</b> `map<string, int32>`<br><b>Description:</b><br>The minimum number of replicas per type of driver.<br> The key is the driver type, the value is the minimum number of replicas.<br> The minimum replicas is guaranteed to be running at all times even if the total number of replicas exceeds the maximum number of replicas set in max_replicas. |
| disableDataProxyProcessing | <b>Type:</b> `bool`<br><b>Description:</b><br>Disable data proxy to process data from ouro temp tables. |

