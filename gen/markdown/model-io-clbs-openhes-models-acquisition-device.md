# Model: io.clbs.openhes.models.acquisition.Device

Defines a device model representing an energy meter or IoT device in the system.
 A device is a physical or logical entity that can be communicated with via one or more communication units,
 has a specific driver type for protocol handling, and belongs to device groups for organizational purposes.

## Fields

| Field | Information |
| --- | --- |
| spec | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceSpec`](model-io-clbs-openhes-models-acquisition-devicespec.md)<br><b>Description:</b><br>The device specification containing configuration and identification details. |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceStatus`](model-io-clbs-openhes-models-acquisition-devicestatus.md)<br><b>Description:</b><br>The current operational status including connection state and statistics. |
| metadata | <b>Type:</b> [`io.clbs.openhes.models.common.MetadataFields`](model-io-clbs-openhes-models-common-metadatafields.md)<br><b>Description:</b><br>Metadata including id, name, generation, user-managed fields, and system-managed fields. |

