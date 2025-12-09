# Model: io.clbs.openhes.models.acquisition.StreamUploadFirmwareImageRequest

Defines a request for retrieving a block of firmware image data for streamed upload.

## Fields

| Field | Information |
| --- | --- |
| firmwareImageId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique firmware image identifier. |
| fileName | <b>Type:</b> `string`<br><b>Description:</b><br>The firmware image file name.<br><b>Example:</b> "firmware.bin" |
| blockSize | <b>Type:</b> `int32`<br><b>Description:</b><br>The size of the block to retrieve in bytes. |
| blockIndex | <b>Type:</b> `int32`<br><b>Description:</b><br>The index of the block to retrieve (0-based). All blocks must be send in order. |
| lastBlock | <b>Type:</b> `bool`<br><b>Description:</b><br>Indicates whether this is the last block in the stream. |
| blockData | <b>Type:</b> `bytes`<br><b>Description:</b><br>The data block of firmware image. |

