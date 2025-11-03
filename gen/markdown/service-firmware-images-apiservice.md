# ApiService - Firmware Images

## CreateFirmwareImage

Creates a new firmware image. Returns the identifier of the newly created firmware image.

```proto
CreateFirmwareImage(io.clbs.openhes.models.acquisition.CreateFirmwareImageRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateFirmwareImageRequest`](model-io-clbs-openhes-models-acquisition-createfirmwareimagerequest.md)
- Output: `google.protobuf.StringValue`

## ListFirmwareImages

Retrieves a paginated list of firmware images based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListFirmwareImages(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfFirmwareImage)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfFirmwareImage`](model-io-clbs-openhes-models-acquisition-listoffirmwareimage.md)

## GetFirmwareImage

Retrieves the details of the specified firmware image.

```proto
GetFirmwareImage(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.FirmwareImage)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.FirmwareImage`](model-io-clbs-openhes-models-acquisition-firmwareimage.md)

## UpdateFirmwareImage

Updates the details of an existing firmware image.

```proto
UpdateFirmwareImage(io.clbs.openhes.models.acquisition.FirmwareImage)
```

- Input: [`io.clbs.openhes.models.acquisition.FirmwareImage`](model-io-clbs-openhes-models-acquisition-firmwareimage.md)

## DeleteFirmwareImage

Deletes the specified firmware image.

```proto
DeleteFirmwareImage(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## StreamUploadFirmwareImageFile

Starts streaming upload of a firmware image file.

```proto
StreamUploadFirmwareImageFile(io.clbs.openhes.models.acquisition.StreamUploadFirmwareImageRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.StreamUploadFirmwareImageRequest`](model-io-clbs-openhes-models-acquisition-streamuploadfirmwareimagerequest.md)

## StreamDownloadFirmwareImageFile

Starts streaming download of a firmware image file.

```proto
StreamDownloadFirmwareImageFile(io.clbs.openhes.models.acquisition.StreamDownloadFirmwareImageFileRequest) returns (io.clbs.openhes.models.acquisition.FirmwareImageBlock)
```

- Input: [`io.clbs.openhes.models.acquisition.StreamDownloadFirmwareImageFileRequest`](model-io-clbs-openhes-models-acquisition-streamdownloadfirmwareimagefilerequest.md)
- Output: [`io.clbs.openhes.models.acquisition.FirmwareImageBlock`](model-io-clbs-openhes-models-acquisition-firmwareimageblock.md)

