# ApiService - Device Data

## GetDeviceData

Retrieves device data of the specified type (register, profile, irregular profile) based on the specified criteria. The method is generic but limited to return

```proto
GetDeviceData(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.DeviceData)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.DeviceData`](model-io-clbs-openhes-models-acquisition-devicedata.md)

## ListDeviceDataInfo

Retrieves a pagianted list of device data info based on the specified criteria. The page size and page number (zero-based) can be defined in the request. The device data info stores various metadata, such as the period of the regular profiles or the timestamp of the last stored value.

```proto
ListDeviceDataInfo(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceDataInfo)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceDataInfo`](model-io-clbs-openhes-models-acquisition-listofdevicedatainfo.md)

## GetDeviceDataRegisters

Retrieves the list of register-type device data based on the specified criteria.

```proto
GetDeviceDataRegisters(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.RegisterValues)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.RegisterValues`](model-io-clbs-openhes-models-acquisition-registervalues.md)

## GetDeviceDataProfiles

Retrieves the list of profile-type device data based on the specified criteria.

```proto
GetDeviceDataProfiles(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.ProfileValues)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.ProfileValues`](model-io-clbs-openhes-models-acquisition-profilevalues.md)

## GetDeviceDataIrregularProfiles

Retrieves the list of irregular profile-type device data based on the specified criteria.

```proto
GetDeviceDataIrregularProfiles(io.clbs.openhes.models.acquisition.GetDeviceDataRequest) returns (io.clbs.openhes.models.acquisition.IrregularProfileValues)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceDataRequest`](model-io-clbs-openhes-models-acquisition-getdevicedatarequest.md)
- Output: [`io.clbs.openhes.models.acquisition.IrregularProfileValues`](model-io-clbs-openhes-models-acquisition-irregularprofilevalues.md)

