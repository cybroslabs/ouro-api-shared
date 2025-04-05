# Model: io.clbs.openhes.models.acquisition.JobActionSet

Sub-message containing job action set specification.
 In comparison the JobAction shall be used only once per bulk but internally it may cover multiple JobActions.
 For example, if the JobActionSet is of the ActionGetRegister type and no variables filter is set
 then then system gets all registers defined for active device configuration template and reads them all.

## Fields

| Field | Information |
| --- | --- |
| variables | <b>Type:</b> `string`<br><b>Description:</b><br>The filter, meaning depends on the action.<br> - GetRegister, GetPeriodicalProfile and GetIrregularProfile: It's the list of variable identifiers, e.g. "A+" defined in the system. If not set then all variables are read.<br> - Others: Not applicable, ignored. |
| getRegister | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetRegister`](model-io-clbs-openhes-models-acquisition-actiongetregister.md)<br><b>Description:</b><br>The get register action specification. |
| getPeriodicalProfile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfile`](model-io-clbs-openhes-models-acquisition-actiongetperiodicalprofile.md)<br><b>Description:</b><br>The get periodical profile action specification. |
| getIrregularProfile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetIrregularProfile`](model-io-clbs-openhes-models-acquisition-actiongetirregularprofile.md)<br><b>Description:</b><br>The get irregular profile action specification. |
| getEvents | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetEvents`](model-io-clbs-openhes-models-acquisition-actiongetevents.md)<br><b>Description:</b><br>The get events action specification. |
| getDeviceInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetDeviceInfo`](model-io-clbs-openhes-models-acquisition-actiongetdeviceinfo.md)<br><b>Description:</b><br>The get device info action specification. |
| syncClock | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSyncClock`](model-io-clbs-openhes-models-acquisition-actionsyncclock.md)<br><b>Description:</b><br>The sync clock action specification. |
| setRelayState | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetRelayState`](model-io-clbs-openhes-models-acquisition-actionsetrelaystate.md)<br><b>Description:</b><br>The set relay state action specification. |
| setDisconnectorState | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetDisconnectorState`](model-io-clbs-openhes-models-acquisition-actionsetdisconnectorstate.md)<br><b>Description:</b><br>The set disconnector state action specification. |
| getTou | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetTou`](model-io-clbs-openhes-models-acquisition-actiongettou.md)<br><b>Description:</b><br>The get tou action specification. |
| setTou | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetTou`](model-io-clbs-openhes-models-acquisition-actionsettou.md)<br><b>Description:</b><br>The set tou action specification. |
| setLimiter | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetLimiter`](model-io-clbs-openhes-models-acquisition-actionsetlimiter.md)<br><b>Description:</b><br>The set limiter action specification. |
| resetBillingPeriod | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionResetBillingPeriod`](model-io-clbs-openhes-models-acquisition-actionresetbillingperiod.md)<br><b>Description:</b><br>The reset billing period action specification. |
| fwUpdate | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionFwUpdate`](model-io-clbs-openhes-models-acquisition-actionfwupdate.md)<br><b>Description:</b><br>The firmware update action specification. |

