# Model: io.clbs.openhes.models.acquisition.JobAction

Defines the job action specification.
 The `JobAction` represents a single action to be performed on a single device.
 For example, if the action is `ActionGetRegister`, it specifies a single register to be read from the devices.

## Fields

| Field | Information |
| --- | --- |
| actionId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier of the action. |
| attributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The action attributes. |
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
| imageTransfer | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ActionFirmwareImageTransfer`](model-io-clbs-openhes-models-acquisition-actionfirmwareimagetransfer.md)<br><b>Description:</b><br>Defines the ActionFirmwareImageTransfer action. |

