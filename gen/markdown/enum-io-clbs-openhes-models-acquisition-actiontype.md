# Enum: io.clbs.openhes.models.acquisition.ActionType

Action types

## Options

| Value | Description |
| --- | --- |
| ACTION_TYPE_GET_REGISTER | Get register value, for example instantaneous values. The action covers specific register. |
| ACTION_TYPE_GET_PERIODICAL_PROFILE | Get periodical profile, for example load-profile. The action covers specific profile column. |
| ACTION_TYPE_GET_IRREGULAR_PROFILE | Get non-periodical profile, for example daily profile or monthly billing registers. The action covers specific profile column and or specific billing register. |
| ACTION_TYPE_GET_EVENTS | Get event log. The action covers specific event log. |
| ACTION_TYPE_GET_DEVICE_INFO | Get device info. The action returns info data about remote device. |
| ACTION_TYPE_SYNC_CLOCK | Synchronize clock. The action synchronizes the time in the device, it can forcefully set it if force attribute is set. |
| ACTION_TYPE_SET_RELAY_STATE | Set relay state. The action covers single relay. |
| ACTION_TYPE_GET_DISCONNECTOR_STATE | Get disconnector state. The action covers single disconnector if multiple disconnectors are present in the device. |
| ACTION_TYPE_SET_DISCONNECTOR_STATE | Set disconnector state. The action covers single relay. The action covers single disconnector if multiple disconnectors are present in the device. |
| ACTION_TYPE_GET_TOU | Get time-of-use table. |
| ACTION_TYPE_SET_TOU | Set time-of-use table. |
| ACTION_TYPE_SET_LIMITER | Set limiter settings. The action covers single limiter. |
| ACTION_TYPE_RESET_BILLING_PERIOD | Reset billing period. |
| ACTION_TYPE_FW_UPDATE | Start firmware update. The action updates starts FW upgrade procedure. |
