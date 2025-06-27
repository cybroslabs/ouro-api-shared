# ApiService - Device Data

## GetDeviceData

The method to returns register/profile/irregular-profile typed device data. The method is generic but limited to return

```proto
GetDeviceData(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.DeviceData)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.DeviceData`](model-io-clbs-openhes-models-acquisition-devicedata.md)

## GetDeviceDataRegisters

The method to stream out register-typed device data.

```proto
GetDeviceDataRegisters(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.RegisterValues)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.RegisterValues`](model-io-clbs-openhes-models-acquisition-registervalues.md)

## GetDeviceDataProfiles

The method to stream out profile-typed device data.

```proto
GetDeviceDataProfiles(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.ProfileValues)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.ProfileValues`](model-io-clbs-openhes-models-acquisition-profilevalues.md)

## GetDeviceDataIrregularProfiles

The method to stream out profile-typed device data.

```proto
GetDeviceDataIrregularProfiles(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.IrregularProfileValues)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.IrregularProfileValues`](model-io-clbs-openhes-models-acquisition-irregularprofilevalues.md)

