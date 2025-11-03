# Model: io.clbs.openhes.models.acquisition.FirmwareImageBlockRequest

Defines a request for retrieving a block of firmware image data for both streamed and non-streamed modes.

## Fields

| Field | Information |
| --- | --- |
| firmwareImageId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique firmware image identifier. |
| fileName | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware image file name. |
| blockSize | <b>Type:</b> `int32`<br><b>Description:</b><br>The size of the block to retrieve in bytes. |
| blockIndex | <b>Type:</b> `int32`<br><b>Description:</b><br>The index of the block to retrieve (0-based). |

