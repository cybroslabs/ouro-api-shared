# ApiService - Time-Of-Use Tables

## CreateTimeOfUseTable

Creates a new time-of-use table. Returns the identifier of the newly created table.

```proto
CreateTimeOfUseTable(io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest`](model-io-clbs-openhes-models-acquisition-createtimeofusetablerequest.md)
- Output: `google.protobuf.StringValue`

## ListTimeOfUseTables

Retrieves a paginated list of time-of-use tables based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListTimeOfUseTables(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable`](model-io-clbs-openhes-models-acquisition-listoftimeofusetable.md)

## GetTimeOfUseTable

Retrieves the details of the spcified time-of-use table.

```proto
GetTimeOfUseTable(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.TimeOfUseTable)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.TimeOfUseTable`](model-io-clbs-openhes-models-acquisition-timeofusetable.md)

## UpdateTimeOfUseTable

Updates the details of an existing time-of-use table.

```proto
UpdateTimeOfUseTable(io.clbs.openhes.models.acquisition.TimeOfUseTable)
```

- Input: [`io.clbs.openhes.models.acquisition.TimeOfUseTable`](model-io-clbs-openhes-models-acquisition-timeofusetable.md)

## DeleteTimeOfUseTable

Deletes the specified time-of-use table.

```proto
DeleteTimeOfUseTable(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

