# Enum: io.clbs.openhes.models.acquisition.ActionType

Defines the supported action types. Each action type corresponds to a specific operation that can be performed on a device.

## Options

| Value | Description |
| --- | --- |
| ACTION_TYPE_GET_REGISTER | Get a register value, for example an instantaneous values. The action covers a specific register. |
| ACTION_TYPE_GET_PERIODICAL_PROFILE | Get a periodical profile, for example a load profile. The action covers a specific profile column. |
| ACTION_TYPE_GET_IRREGULAR_PROFILE | Get a non-periodical profile, for example a daily profile or monthly billing registers. The action covers a specific profile column and/or billing register. |
| ACTION_TYPE_GET_EVENTS | Get an event log. The action covers a specific event log. |
| ACTION_TYPE_GET_DEVICE_INFO | Get device information. The action returns detailed information data about the remote device. |
| ACTION_TYPE_SYNC_CLOCK | Synchronize the device clock. The action sets the time and can forcefully update it if `force` attribute is set. |
| ACTION_TYPE_SET_RELAY_STATE | Set relay state. The action covers a single relay. |
| ACTION_TYPE_GET_DISCONNECTOR_STATE | Get the disconnector state. |
| ACTION_TYPE_SET_DISCONNECTOR_STATE | Set the disconnector state. |
| ACTION_TYPE_GET_TOU | Get time-of-use table. |
| ACTION_TYPE_SET_TOU | Set time-of-use table. |
| ACTION_TYPE_SET_LIMITER | Set limiter settings. The action covers single limiter. |
| ACTION_TYPE_RESET_BILLING_PERIOD | Reset the billing period. |
| ACTION_TYPE_FW_UPDATE | Start a firmware update. |
