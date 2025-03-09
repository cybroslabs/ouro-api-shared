# Model: io.clbs.openhes.models.acquisition.JobAction

Sub-message containing job action specification

## Fields

| Field | Type | Description |
| --- | --- | --- |
| actionId | string | The action identifier. |
| attributes | map<string, io.clbs.openhes.models.common.FieldValue> | The action attributes. |
| getRegister | io.clbs.openhes.models.acquisition.ActionGetRegister | The get register action specification. |
| getPeriodicalProfile | io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfile | The get periodical profile action specification. |
| getIrregularProfile | io.clbs.openhes.models.acquisition.ActionGetIrregularProfile | The get irregular profile action specification. |
| getEvents | io.clbs.openhes.models.acquisition.ActionGetEvents | The get events action specification. |
| getDeviceInfo | io.clbs.openhes.models.acquisition.ActionGetDeviceInfo | The get device info action specification. |
| syncClock | io.clbs.openhes.models.acquisition.ActionSyncClock | The sync clock action specification. |
| setRelayState | io.clbs.openhes.models.acquisition.ActionSetRelayState | The set relay state action specification. |
| setDisconnectorState | io.clbs.openhes.models.acquisition.ActionSetDisconnectorState | The set disconnector state action specification. |
| getTou | io.clbs.openhes.models.acquisition.ActionGetTou | The get tou action specification. |
| setTou | io.clbs.openhes.models.acquisition.ActionSetTou | The set tou action specification. |
| setLimiter | io.clbs.openhes.models.acquisition.ActionSetLimiter | The set limiter action specification. |
| resetBillingPeriod | io.clbs.openhes.models.acquisition.ActionResetBillingPeriod | The reset billing period action specification. |
| fwUpdate | io.clbs.openhes.models.acquisition.ActionFwUpdate | The firmware update action specification. |

