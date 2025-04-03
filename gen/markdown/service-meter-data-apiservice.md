# ApiService - Meter Data

## GetMeterDataRegisters

The method to stream out register-typed meter data.

```proto
GetMeterDataRegisters(io.clbs.openhes.models.acquisition.GetMeterDataRequest) returns (io.clbs.openhes.models.acquisition.RegisterValues)
```

- Input: [io.clbs.openhes.models.acquisition.GetMeterDataRequest](model-io-clbs-openhes-models-acquisition-getmeterdatarequest.md)
- Output: [io.clbs.openhes.models.acquisition.RegisterValues](model-io-clbs-openhes-models-acquisition-registervalues.md)

## GetMeterDataProfiles

The method to stream out profile-typed meter data.

```proto
GetMeterDataProfiles(io.clbs.openhes.models.acquisition.GetMeterDataRequest) returns (io.clbs.openhes.models.acquisition.ProfileValues)
```

- Input: [io.clbs.openhes.models.acquisition.GetMeterDataRequest](model-io-clbs-openhes-models-acquisition-getmeterdatarequest.md)
- Output: [io.clbs.openhes.models.acquisition.ProfileValues](model-io-clbs-openhes-models-acquisition-profilevalues.md)

## GetMeterDataIrregularProfiles

The method to stream out profile-typed meter data.

```proto
GetMeterDataIrregularProfiles(io.clbs.openhes.models.acquisition.GetMeterDataRequest) returns (io.clbs.openhes.models.acquisition.IrregularProfileValues)
```

- Input: [io.clbs.openhes.models.acquisition.GetMeterDataRequest](model-io-clbs-openhes-models-acquisition-getmeterdatarequest.md)
- Output: [io.clbs.openhes.models.acquisition.IrregularProfileValues](model-io-clbs-openhes-models-acquisition-irregularprofilevalues.md)

