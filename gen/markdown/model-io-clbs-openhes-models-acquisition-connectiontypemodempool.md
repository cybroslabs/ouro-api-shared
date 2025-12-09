# Model: io.clbs.openhes.models.acquisition.ConnectionTypeModemPool

Defines the connection information for a phone line (modem) connection type.

## Fields

| Field | Information |
| --- | --- |
| number | <b>Type:</b> `string`<br><b>Description:</b><br>The phone number of the device to connect to. |
| poolId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique modem pool identifier. A modem pool is a group of modems that can be used to connect to devices. The final modem is selected by the Taskmaster at job start time. |
| modem | <b>Type:</b> [`io.clbs.openhes.models.acquisition.ModemInfo`](model-io-clbs-openhes-models-acquisition-modeminfo.md)<br><b>Description:</b><br>The modem device assigned to the job. This field is filled only when the connection type is modem. The value is assigned by the Taskmaster at the start of the job, and the driver must use this modem exclusively for this job. |

