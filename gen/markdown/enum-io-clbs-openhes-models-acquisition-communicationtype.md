# Enum: io.clbs.openhes.models.acquisition.CommunicationType

Defines the supported communication types for device connections.

## Options

| Value | Description |
| --- | --- |
| COMMUNICATION_TYPE_UNSPECIFIED | Unspecified communication type. |
| COMMUNICATION_TYPE_TCPIP | The communication is established via direct TCP/IP. |
| COMMUNICATION_TYPE_MODEM_POOL | The communication is established via phone line (modem). |
| COMMUNICATION_TYPE_SERIAL_LINE_DIRECT | The communication is established via a direct serial line. |
| COMMUNICATION_TYPE_SERIAL_LINE_MOXA | The communication is established via a controlled  serial line (Moxa). |
| COMMUNICATION_TYPE_LISTENING | The communication is passive. The driver listens on a port (for example, DLMS devices in push mode) or subscribes to a message queue (for example, MQTT). |
