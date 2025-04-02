# Model: io.clbs.openhes.models.acquisition.DeviceSpec

Sub-message - the device specification.

## Fields

| Field | Information |
| --- | --- |
| dctId | <b>Type:</b> `string`<br><b>Description:</b><br>@gqltype: UUID<br><br>The device configuration template identifier. |
| externalId | <b>Type:</b> `string`<br><b>Description:</b><br>The external identifier of the device. |
| attributes | <b>Type:</b> `map<string, [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md)>`<br><b>Description:</b><br>The template of the action attributes. It is represented as a list of attribute definitions. |
| communicationUnitLink | <b>Type:</b> `[io.clbs.openhes.models.acquisition.DeviceCommunicationUnit](model-io-clbs-openhes-models-acquisition-devicecommunicationunit.md)`<br><b>Description:</b><br>The list of communication unit identifiers (and additional info) that set CUs usable to communicate with the device. It's an ordered set where the first element is the primary communication unit with the highest priority. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone related to the device, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2". |

