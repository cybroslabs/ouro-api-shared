# Model: io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest

Defines the specification for removing device configuration registers to device configuration templates.

## Fields

| Field | Information |
| --- | --- |
| dctId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique device configuration template identifier. |
| registerId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>A list of device configuration register identifiers used by standard readout. |
| scadaRegisterId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>A list of device configuration register identifiers used by SCADA readout. |

