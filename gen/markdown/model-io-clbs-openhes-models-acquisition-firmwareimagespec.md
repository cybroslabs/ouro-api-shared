# Model: io.clbs.openhes.models.acquisition.FirmwareImageSpec

Defines a firmware image specification.

## Fields

| Field | Information |
| --- | --- |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type. |
| version | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware version. |
| content | <b>Type:</b> [`io.clbs.openhes.models.acquisition.FirmwareImageFile`](model-io-clbs-openhes-models-acquisition-firmwareimagefile.md)<br><b>Description:</b><br>The list of firmware image files. |
| description | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware image description. |
| attributes | <b>Type:</b> map<`string`, [`io.clbs.openhes.models.common.FieldValue`](model-io-clbs-openhes-models-common-fieldvalue.md)><br><b>Description:</b><br>The firmware attributes, depending on the driver type. |

