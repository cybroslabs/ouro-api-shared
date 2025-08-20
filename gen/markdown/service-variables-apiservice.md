# ApiService - Variables

## CreateVariable

Creates a new variable. A variable is a named abstraction for device configuration registers. Returns the identifier of the newly created variable.

```proto
CreateVariable(io.clbs.openhes.models.acquisition.CreateVariableRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateVariableRequest`](model-io-clbs-openhes-models-acquisition-createvariablerequest.md)
- Output: `google.protobuf.StringValue`

## ListVariables

Retrieves a paginated list of variables based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListVariables(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfVariable)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfVariable`](model-io-clbs-openhes-models-acquisition-listofvariable.md)

## GetVariable

Retrieves the details of the specified variable.

```proto
GetVariable(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Variable)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.Variable`](model-io-clbs-openhes-models-acquisition-variable.md)

## UpdateVariable

Updates the details of an existing variable. Fields that are omitted from the request will be left unchanged.

```proto
UpdateVariable(io.clbs.openhes.models.acquisition.Variable)
```

- Input: [`io.clbs.openhes.models.acquisition.Variable`](model-io-clbs-openhes-models-acquisition-variable.md)

## DeleteVariable

Deletes the specified variable.

```proto
DeleteVariable(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## AddRegisterToVariable

Adds a specified register to and existing variable.

```proto
AddRegisterToVariable(io.clbs.openhes.models.acquisition.AddRegisterToVariableRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.AddRegisterToVariableRequest`](model-io-clbs-openhes-models-acquisition-addregistertovariablerequest.md)

## RemoveRegisterFromVariable

Removes a specified register from a variable.

```proto
RemoveRegisterFromVariable(io.clbs.openhes.models.acquisition.RemoveRegisterFromVariableRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.RemoveRegisterFromVariableRequest`](model-io-clbs-openhes-models-acquisition-removeregisterfromvariablerequest.md)

