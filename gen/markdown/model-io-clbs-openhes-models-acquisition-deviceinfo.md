# Model: io.clbs.openhes.models.acquisition.DeviceInfo

Message holds common device information.

## Fields

| Field | Information |
| --- | --- |
| infoTimestamp | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the values were read. |
| manufacturerSerialNumber | <b>Type:</b> `string`<br><b>Description:</b><br>The serial number of the device set by manufacturer. Typical source: 0-0:42.0.0.255 |
| deviceSerialNumber | <b>Type:</b> `string`<br><b>Description:</b><br>The device serial number of the device.  Typical source: 0-0:96.1.0.255 |
| firmwareVersion | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware version identifier of the device. Typical source: 0-0:0.2.0.255 |
| clockDelta | <b>Type:</b> `google.protobuf.Duration`<br><b>Description:</b><br>The clock delta against acquisition server clock (got by a driver) where positive value means that the device clock is ahead of local clock and vice versa. |
| deviceModel | <b>Type:</b> `string`<br><b>Description:</b><br>The model of the device. |
| errorRegister | <b>Type:</b> `uint64`<br><b>Description:</b><br>The error register 0. Typical source: 0-0:97.97.1.255 (higher 32-bit) + 0-0:97.97.0.255 (lower 32-bit) |
| relayStates | <b>Type:</b> `bool`<br><b>Description:</b><br>The list of relay states. The order of the relays is the same as in the device. The value is true if the relay is connected and false if the relay is disconnected. |
| connectionState | <b>Type:</b> `bool`<br><b>Description:</b><br>The state of the connection. It represents the disconnector state where true means that the customer is connected and false means that the customer is disconnected. |

