# ApiService - Device Configuration Templates

## CreateDeviceConfigurationTemplate

Creates a new device configuration template. Returns the identifier of the newly created template.

```proto
CreateDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest`](model-io-clbs-openhes-models-acquisition-createdeviceconfigurationtemplaterequest.md)
- Output: `google.protobuf.StringValue`

## ListDeviceConfigurationTemplates

Retrieves a paginated list of device configuration templates based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListDeviceConfigurationTemplates(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate`](model-io-clbs-openhes-models-acquisition-listofdeviceconfigurationtemplate.md)

## GetDeviceConfigurationTemplate

Retrieves the details of the specified device configuration template.

```proto
GetDeviceConfigurationTemplate(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate`](model-io-clbs-openhes-models-acquisition-deviceconfigurationtemplate.md)

## UpdateDeviceConfigurationTemplate

Updates the details of an existing device configuration template. Fields that are omitted from the request will be left unchanged.

```proto
UpdateDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate)
```

- Input: [`io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate`](model-io-clbs-openhes-models-acquisition-deviceconfigurationtemplate.md)

## DeleteDeviceConfigurationTemplate

Deletes the specified device configuration template.

```proto
DeleteDeviceConfigurationTemplate(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## AddDeviceConfigurationRegisterToDeviceConfigurationTemplate

Adds a specified device configuration register to an existing device configuration template.

```proto
AddDeviceConfigurationRegisterToDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest`](model-io-clbs-openhes-models-acquisition-adddeviceconfigurationregistertodeviceconfigurationtemplaterequest.md)

## RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate

Removes a specified device configuration register from a device configuration template.

```proto
RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest`](model-io-clbs-openhes-models-acquisition-removedeviceconfigurationregisterfromdeviceconfigurationtemplaterequest.md)

