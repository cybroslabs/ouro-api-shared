# Model: io.clbs.openhes.models.acquisition.DeviceInfo

Message holds common device information.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| infoTimestamp | google.protobuf.Timestamp | The timestamp when the values were read. |
| manufacturerSerialNumber | string | The serial number of the device set by manufacturer. Typical source: 0-0:42.0.0.255 |
| deviceSerialNumber | string | The device serial number of the device.  Typical source: 0-0:96.1.0.255 |
| firmwareVersion | string | The firmware version identifier of the device. Typical source: 0-0:0.2.0.255 |
| clockDelta | google.protobuf.Duration | The clock delta against acquisition server clock (got by a driver) where positive value means that the device clock is ahead of local clock and vice versa. |
| deviceModel | string | The model of the device. |
| errorRegister | uint64 | The error register 0. Typical source: 0-0:97.97.1.255 (higher 32-bit) + 0-0:97.97.0.255 (lower 32-bit) |
| relayStates | bool | The list of relay states. The order of the relays is the same as in the device. The value is true if the relay is connected and false if the relay is disconnected. |
| connectionState | bool | The state of the connection. It represents the disconnector state where true means that the customer is connected and false means that the customer is disconnected. |

