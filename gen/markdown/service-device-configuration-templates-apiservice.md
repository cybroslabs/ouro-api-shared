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

## AddRegisterToDeviceConfigurationTemplate

```proto
AddRegisterToDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.AddRegisterToDeviceConfigurationTemplateRequest)
```

- Input: [io.clbs.openhes.models.acquisition.AddRegisterToDeviceConfigurationTemplateRequest](model-io-clbs-openhes-models-acquisition-addregistertodeviceconfigurationtemplaterequest.md)

## RemoveRegisterFromDeviceConfigurationTemplate

```proto
RemoveRegisterFromDeviceConfigurationTemplate(io.clbs.openhes.models.acquisition.RemoveRegisterFromDeviceConfigurationTemplateRequest)
```

- Input: [io.clbs.openhes.models.acquisition.RemoveRegisterFromDeviceConfigurationTemplateRequest](model-io-clbs-openhes-models-acquisition-removeregisterfromdeviceconfigurationtemplaterequest.md)

