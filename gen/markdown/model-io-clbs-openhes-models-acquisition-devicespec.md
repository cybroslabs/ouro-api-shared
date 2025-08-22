# Model: io.clbs.openhes.models.acquisition.DeviceSpec

Defines the device specification.

## Fields

| Field | Information |
| --- | --- |
| dctId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The identifier of the device configuration template. |
| externalId | <b>Type:</b> `string`<br><b>Description:</b><br>The external identifier of the device. |
| communicationUnitLink | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceCommunicationUnit`](model-io-clbs-openhes-models-acquisition-devicecommunicationunit.md)<br><b>Description:</b><br>The ordered list of communication units (with additional info) linked to the device. The first entry is the primary communication unit with the highest priority. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone associated with the device. For example, `"America/New_York"`, `"Europe/Prague"`, `"CET"`, `"GMT"`, `"Etc/GMT+2"`. |
| defaultLinkAttributes | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DefaultDeviceCommunicationUnitAttributes`](model-io-clbs-openhes-models-acquisition-defaultdevicecommunicationunitattributes.md)<br><b>Description:</b><br>The default communication unit link attributes used when the device is created without a communication unit link. Typically usable for floating devices that are connected via data concentrators and related driver assigns the communication unit dynamically. The list must not contain same application protocol more than once. |

