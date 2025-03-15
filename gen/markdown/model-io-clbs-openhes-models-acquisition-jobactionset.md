# Model: io.clbs.openhes.models.acquisition.JobActionSet

Sub-message containing job action set specification.
 In comparison the JobAction shall be used only once per bulk but internally it may cover multiple JobActions.
 For example, if the JobActionSet is of the ActionGetRegister type and no variables filter is set
 then then system gets all registers defined for active device configuration template and reads them all.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| variables | string | The filter, meaning depends on the action.<br> - GetRegister, GetPeriodicalProfile and GetIrregularProfile: It's the list of variable identifiers, e.g. "A+" defined in the system. If not set then all variables are read.<br> - Others: Not applicable, ignored. |
| getRegister | [io.clbs.openhes.models.acquisition.ActionGetRegister](model-io-clbs-openhes-models-acquisition-actiongetregister.md) | The get register action specification. |
| getPeriodicalProfile | [io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfile](model-io-clbs-openhes-models-acquisition-actiongetperiodicalprofile.md) | The get periodical profile action specification. |
| getIrregularProfile | [io.clbs.openhes.models.acquisition.ActionGetIrregularProfile](model-io-clbs-openhes-models-acquisition-actiongetirregularprofile.md) | The get irregular profile action specification. |
| getEvents | [io.clbs.openhes.models.acquisition.ActionGetEvents](model-io-clbs-openhes-models-acquisition-actiongetevents.md) | The get events action specification. |
| getDeviceInfo | [io.clbs.openhes.models.acquisition.ActionGetDeviceInfo](model-io-clbs-openhes-models-acquisition-actiongetdeviceinfo.md) | The get device info action specification. |
| syncClock | [io.clbs.openhes.models.acquisition.ActionSyncClock](model-io-clbs-openhes-models-acquisition-actionsyncclock.md) | The sync clock action specification. |
| setRelayState | [io.clbs.openhes.models.acquisition.ActionSetRelayState](model-io-clbs-openhes-models-acquisition-actionsetrelaystate.md) | The set relay state action specification. |
| setDisconnectorState | [io.clbs.openhes.models.acquisition.ActionSetDisconnectorState](model-io-clbs-openhes-models-acquisition-actionsetdisconnectorstate.md) | The set disconnector state action specification. |
| getTou | [io.clbs.openhes.models.acquisition.ActionGetTou](model-io-clbs-openhes-models-acquisition-actiongettou.md) | The get tou action specification. |
| setTou | [io.clbs.openhes.models.acquisition.ActionSetTou](model-io-clbs-openhes-models-acquisition-actionsettou.md) | The set tou action specification. |
| setLimiter | [io.clbs.openhes.models.acquisition.ActionSetLimiter](model-io-clbs-openhes-models-acquisition-actionsetlimiter.md) | The set limiter action specification. |
| resetBillingPeriod | [io.clbs.openhes.models.acquisition.ActionResetBillingPeriod](model-io-clbs-openhes-models-acquisition-actionresetbillingperiod.md) | The reset billing period action specification. |
| fwUpdate | [io.clbs.openhes.models.acquisition.ActionFwUpdate](model-io-clbs-openhes-models-acquisition-actionfwupdate.md) | The firmware update action specification. |

