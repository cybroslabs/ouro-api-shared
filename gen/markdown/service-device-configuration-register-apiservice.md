# ApiService - Device Configuration Register

## CreateDeviceConfigurationRegister

Creates a new register. The register object holds the information about the single device register.

```proto
CreateDeviceConfigurationRegister(io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest`](model-io-clbs-openhes-models-acquisition-createdeviceconfigurationregisterrequest.md)
- Output: `google.protobuf.StringValue`

## ListDeviceConfigurationRegisters

```proto
ListDeviceConfigurationRegisters(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister`](model-io-clbs-openhes-models-acquisition-listofdeviceconfigurationregister.md)

## GetDeviceConfigurationRegister

```proto
GetDeviceConfigurationRegister(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceConfigurationRegister)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.DeviceConfigurationRegister`](model-io-clbs-openhes-models-acquisition-deviceconfigurationregister.md)

## UpdateDeviceConfigurationRegister

```proto
UpdateDeviceConfigurationRegister(io.clbs.openhes.models.acquisition.DeviceConfigurationRegister)
```

- Input: [`io.clbs.openhes.models.acquisition.DeviceConfigurationRegister`](model-io-clbs-openhes-models-acquisition-deviceconfigurationregister.md)

## DeleteDeviceConfigurationRegister

```proto
DeleteDeviceConfigurationRegister(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

