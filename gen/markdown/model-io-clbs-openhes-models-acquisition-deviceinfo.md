# Model: io.clbs.openhes.models.acquisition.DeviceInfo

Defines common device information.

## Fields

| Field | Information |
| --- | --- |
| infoTimestamp | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the values were read. |
| manufacturerSerialNumber | <b>Type:</b> `string`<br><b>Description:</b><br>The serial number of the device assigned by the manufacturer. Typical source: `0-0:42.0.0.255`. |
| deviceSerialNumber | <b>Type:</b> `string`<br><b>Description:</b><br>The device serial number.  Typical source: `0-0:96.1.0.255`. |
| firmwareVersion | <b>Type:</b> `string`<br><b>Description:</b><br>The device firmware version identifier. Typical source: `0-0:0.2.0.255`. |
| clockDelta | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The time difference (delta) between the device clock and the acquisition server clock (driver). Positive values mean the device clock is ahead, negative values mean the device clock is behind. |
| deviceModel | <b>Type:</b> `string`<br><b>Description:</b><br>The device model. |
| errorRegister | <b>Type:</b> `uint64`<br><b>Description:</b><br>The error register (register `0`). Typical source: `0-0:97.97.1.255` (higher 32-bit) + `0-0:97.97.0.255` (lower 32-bit). |
| relayStates | <b>Type:</b> `bool`<br><b>Description:</b><br>The list of relay states. The order of the relays is the same as in the device. The value is `true` if the relay is connected and `false` if the relay is disconnected. |
| connectionState | <b>Type:</b> `bool`<br><b>Description:</b><br>The connection (disconnector) state. The value is `true` if the customer is connected and `false` if the customer is disconnected. |

