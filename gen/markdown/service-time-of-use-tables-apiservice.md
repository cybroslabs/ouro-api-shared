# ApiService - Time-Of-Use Tables

## CreateTimeOfUseTable

The method to create a new time-of-use table.

```proto
CreateTimeOfUseTable(io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest`](model-io-clbs-openhes-models-acquisition-createtimeofusetablerequest.md)
- Output: `google.protobuf.StringValue`

## ListTimeOfUseTables

The method to get the list of time-of-use tables.

```proto
ListTimeOfUseTables(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable`](model-io-clbs-openhes-models-acquisition-listoftimeofusetable.md)

## GetTimeOfUseTable

The method to get the time-of-use table.

```proto
GetTimeOfUseTable(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.TimeOfUseTable)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.TimeOfUseTable`](model-io-clbs-openhes-models-acquisition-timeofusetable.md)

## UpdateTimeOfUseTable

The method to update the time-of-use table.

```proto
UpdateTimeOfUseTable(io.clbs.openhes.models.acquisition.TimeOfUseTable)
```

- Input: [`io.clbs.openhes.models.acquisition.TimeOfUseTable`](model-io-clbs-openhes-models-acquisition-timeofusetable.md)

## DeleteTimeOfUseTable

The method to delete the time-of-use table.

```proto
DeleteTimeOfUseTable(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

