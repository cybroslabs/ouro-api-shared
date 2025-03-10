# ApiService - Device Registers

## CreateDeviceRegister

Creates a new register. The register object holds the information about the single device register.

```proto
CreateDeviceRegister(io.clbs.openhes.models.acquisition.CreateDeviceRegisterRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.CreateDeviceRegisterRequest](model-io-clbs-openhes-models-acquisition-createdeviceregisterrequest.md)
- Output: google.protobuf.StringValue

## ListDeviceRegisters

```proto
ListDeviceRegisters(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceRegister)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfDeviceRegister](model-io-clbs-openhes-models-acquisition-listofdeviceregister.md)

## GetDeviceRegister

```proto
GetDeviceRegister(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceRegister)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.DeviceRegister](model-io-clbs-openhes-models-acquisition-deviceregister.md)

## UpdateDeviceRegister

```proto
UpdateDeviceRegister(io.clbs.openhes.models.acquisition.DeviceRegister)
```

- Input: [io.clbs.openhes.models.acquisition.DeviceRegister](model-io-clbs-openhes-models-acquisition-deviceregister.md)

## DeleteDeviceRegister

```proto
DeleteDeviceRegister(google.protobuf.StringValue)
```

- Input: google.protobuf.StringValue

