# Model: io.clbs.openhes.models.acquisition.DeviceSpec

Sub-message - the device specification.

## Fields

| Field | Information |
| --- | --- |
| dctId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The device configuration template identifier. |
| externalId | <b>Type:</b> `string`<br><b>Description:</b><br>The external identifier of the device. |
| communicationUnitLink | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DeviceCommunicationUnit`](model-io-clbs-openhes-models-acquisition-devicecommunicationunit.md)<br><b>Description:</b><br>The list of communication unit identifiers (and additional info) that set CUs usable to communicate with the device. It's an ordered set where the first element is the primary communication unit with the highest priority. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone related to the device, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2". |

