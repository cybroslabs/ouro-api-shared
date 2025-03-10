# Model: io.clbs.openhes.models.acquisition.DeviceSpec

Sub-message - the device specification.

## Fields

| Field | Type | Description |
| --- | --- | --- |
| driverType | string | The driver type identifier. |
| externalId | string | The external identifier of the device. |
| attributes | map<string, [io.clbs.openhes.models.common.FieldValue](model-io-clbs-openhes-models-common-fieldvalue.md)> | The template of the action attributes. It is represented as a list of attribute definitions. |
| communicationUnitLink | [io.clbs.openhes.models.acquisition.DeviceCommunicationUnit](model-io-clbs-openhes-models-acquisition-devicecommunicationunit.md) | The list of communication unit identifiers (and additional info) that set CUs usable to communicate with the device. It's an ordered set where the first element is the primary communication unit with the highest priority. |
| timezone | string | The timezone related to the device, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2". |

