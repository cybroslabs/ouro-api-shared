# ApiService - Driver Info

## ListDrivers

Retrieves the list of drivers.

```proto
ListDrivers(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDriver)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfDriver](model-io-clbs-openhes-models-acquisition-listofdriver.md)

## GetDriver

Retrieves the driver.

```proto
GetDriver(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Driver)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.Driver](model-io-clbs-openhes-models-acquisition-driver.md)

