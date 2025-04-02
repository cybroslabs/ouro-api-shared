# Model: io.clbs.openhes.models.acquisition.ConnectionTypeModemPool

Sub-message containing connection info for phone line (modem) connection type

## Fields

| Field | Information |
| --- | --- |
| number | <b>Type:</b> `string`<br><b>Description:</b><br>The phone number of the device to connect to. |
| poolId | <b>Type:</b> `string`<br><b>Description:</b><br>The modem pool identifier. The pool is a group of modems that can be used to connect to the device. Final modem is selected by the Taskmaster at the time of the job start. |
| modem | <b>Type:</b> `[io.clbs.openhes.models.acquisition.ModemInfo](model-io-clbs-openhes-models-acquisition-modeminfo.md)`<br><b>Description:</b><br>The modem device assigned to the job. This is filled only and only when the connection type is modem. The value is assigned by the Taskmaster when to job is being started. Driver is required to use this modem device to connect to the meter only and only for the time of this job! |

