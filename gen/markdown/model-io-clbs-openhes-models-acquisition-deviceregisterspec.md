# Model: io.clbs.openhes.models.acquisition.DeviceRegisterSpec

## Fields

| Field | Type | Description |
| --- | --- | --- |
| registerType | string | The technical/internal ID of the register. |
| displayName | string | The display name of the register. Must be in format '<manufacturer> <device_type> [<device_type_version>]'.<br> It must respect upper/lower characters.<br> The generic registers, such as 'cybros labs generic', must be named as '<driver_company_name> generic'.<br><br> Examples: 'Addax NP73E', 'cybros labs generic', 'Landis+Gyr S650 v2' |

