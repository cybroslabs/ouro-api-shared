# ApiService - Driver Info

## ListDrivers

Retrieves a paginated list of drivers based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListDrivers(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDriver)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDriver`](model-io-clbs-openhes-models-acquisition-listofdriver.md)

## GetDriver

Retrieves the details of the specified driver.

```proto
GetDriver(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Driver)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.Driver`](model-io-clbs-openhes-models-acquisition-driver.md)

