# Model: io.clbs.openhes.models.acquisition.FirmwareImageSpec

Defines a firmware image specification.

## Fields

| Field | Information |
| --- | --- |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type.<br><b>Example:</b> "LANDISGYR_E650_DLMS_SN" |
| version | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware version.<br><b>Example:</b> "v2.0.1" |
| content | <b>Type:</b> [`io.clbs.openhes.models.acquisition.FirmwareImageFile`](model-io-clbs-openhes-models-acquisition-firmwareimagefile.md)<br><b>Description:</b><br>The list of firmware image files. |
| description | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware image description.<br><b>Example:</b> "Firmware update for ST402D meters - bug fixes and performance improvements" |
| attributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The firmware attributes, depending on the driver type. |

