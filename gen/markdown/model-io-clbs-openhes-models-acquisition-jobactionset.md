# Model: io.clbs.openhes.models.acquisition.JobActionSet

Defines the job action set specification.
 Unlike a single `JobAction` that is used only once per bulk. `JobActionSet` may internally cover multiple `JobActions`.
 For example, if the action type is `GetRegister` and no variable filter is specified, the system automatically retrieves all registers defined in the active device configuration template.

## Fields

| Field | Information |
| --- | --- |
| variables | <b>Type:</b> `string`<br><b>Description:</b><br>The variable filter. Meaning depends on the action type:<br>  - `GetRegister`, `GetPeriodicalProfile` and `GetIrregularProfile`: List of variable identifiers (for example, `"A+"`) defined in the system. If not set, all variables are read.<br> - Others: Not applicable (ignored). |
| getRegister | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetRegister`](model-io-clbs-openhes-models-acquisition-actiongetregister.md)<br><b>Description:</b><br>Defines the GetRegister action. |
| getPeriodicalProfile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfile`](model-io-clbs-openhes-models-acquisition-actiongetperiodicalprofile.md)<br><b>Description:</b><br>Defines the GetPeriodicalProfile action. |
| getIrregularProfile | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetIrregularProfile`](model-io-clbs-openhes-models-acquisition-actiongetirregularprofile.md)<br><b>Description:</b><br>Defines the GetIrregularProfile action. |
| getEvents | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetEvents`](model-io-clbs-openhes-models-acquisition-actiongetevents.md)<br><b>Description:</b><br>Defines the GetEvents action. |
| getDeviceInfo | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetDeviceInfo`](model-io-clbs-openhes-models-acquisition-actiongetdeviceinfo.md)<br><b>Description:</b><br>Defines the GetDeviceInfo action. |
| syncClock | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSyncClock`](model-io-clbs-openhes-models-acquisition-actionsyncclock.md)<br><b>Description:</b><br>Defines the SyncClock action. |
| setRelayState | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetRelayState`](model-io-clbs-openhes-models-acquisition-actionsetrelaystate.md)<br><b>Description:</b><br>Defines the SetRelayState action. |
| setDisconnectorState | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetDisconnectorState`](model-io-clbs-openhes-models-acquisition-actionsetdisconnectorstate.md)<br><b>Description:</b><br>Defines the SetDisconnectorState action. |
| getTou | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionGetTou`](model-io-clbs-openhes-models-acquisition-actiongettou.md)<br><b>Description:</b><br>Defines the GetTou action. |
| setTou | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetTou`](model-io-clbs-openhes-models-acquisition-actionsettou.md)<br><b>Description:</b><br>Defines the SetTou action.. |
| setLimiter | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionSetLimiter`](model-io-clbs-openhes-models-acquisition-actionsetlimiter.md)<br><b>Description:</b><br>Defines the SetLimiter action. |
| resetBillingPeriod | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionResetBillingPeriod`](model-io-clbs-openhes-models-acquisition-actionresetbillingperiod.md)<br><b>Description:</b><br>Defines the ResetBillingPeriod action. |
| fwUpdate | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionFwUpdate`](model-io-clbs-openhes-models-acquisition-actionfwupdate.md)<br><b>Description:</b><br>Defines the firmware update action. |

