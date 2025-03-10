# ApiService - Device Configuration Templates

## CreateDeviceConfigurationTemplate

```proto
CreateDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest](model-io-clbs-openhes-models-acquisition-createdeviceconfigurationtemplaterequest.md)
- Output: google.protobuf.StringValue

## ListDeviceConfigurationTemplates

```proto
ListDeviceConfigurationTemplates(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate](model-io-clbs-openhes-models-acquisition-listofdeviceconfigurationtemplate.md)

## GetDeviceConfigurationTemplate

```proto
GetDeviceConfigurationTemplate(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate](model-io-clbs-openhes-models-acquisition-deviceconfigurationtemplate.md)

## UpdateDeviceConfigurationTemplate

```proto
UpdateDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate)
```

- Input: [io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate](model-io-clbs-openhes-models-acquisition-deviceconfigurationtemplate.md)

## DeleteDeviceConfigurationTemplate

```proto
DeleteDeviceConfigurationTemplate(google.protobuf.StringValue)
```

- Input: google.protobuf.StringValue

## AddDeviceConfigurationRegisterToDeviceConfigurationTemplate

```proto
AddDeviceConfigurationRegisterToDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest)
```

- Input: [io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest](model-io-clbs-openhes-models-acquisition-adddeviceconfigurationregistertodeviceconfigurationtemplaterequest.md)

## RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate

```proto
RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest)
```

- Input: [io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest](model-io-clbs-openhes-models-acquisition-removedeviceconfigurationregisterfromdeviceconfigurationtemplaterequest.md)

