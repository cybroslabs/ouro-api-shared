# ApiService - Device Configuration Registers

## CreateDeviceConfigurationRegister

Creates a new device configuration register. Returns the identifier of the newly created register.

```proto
CreateDeviceConfigurationRegister(io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest`](model-io-clbs-openhes-models-acquisition-createdeviceconfigurationregisterrequest.md)
- Output: `google.protobuf.StringValue`

## ListDeviceConfigurationRegisters

Retrieves a paginated list of configuration registers based on the specified criteria. The page size and page number (zero-based) are defined in the request.

```proto
ListDeviceConfigurationRegisters(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister`](model-io-clbs-openhes-models-acquisition-listofdeviceconfigurationregister.md)

## GetDeviceConfigurationRegister

Retrieves the details of the specified device configuration register.

```proto
GetDeviceConfigurationRegister(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceConfigurationRegister)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.DeviceConfigurationRegister`](model-io-clbs-openhes-models-acquisition-deviceconfigurationregister.md)

## UpdateDeviceConfigurationRegister

Updates the details of an existing device configuration register. Fields that are omitted from the request will be left unchanged.

```proto
UpdateDeviceConfigurationRegister(io.clbs.openhes.models.acquisition.DeviceConfigurationRegister)
```

- Input: [`io.clbs.openhes.models.acquisition.DeviceConfigurationRegister`](model-io-clbs-openhes-models-acquisition-deviceconfigurationregister.md)

## DeleteDeviceConfigurationRegister

Deletes the specified device configuration register.

```proto
DeleteDeviceConfigurationRegister(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

