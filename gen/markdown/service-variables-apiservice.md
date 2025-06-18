# ApiService - Variables

## CreateVariable

Creates a new variable. The variable object defines named variable that provides abstraction for device configuration registers.

```proto
CreateVariable(io.clbs.openhes.models.acquisition.CreateVariableRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateVariableRequest`](model-io-clbs-openhes-models-acquisition-createvariablerequest.md)
- Output: `google.protobuf.StringValue`

## ListVariables

```proto
ListVariables(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfVariable)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfVariable`](model-io-clbs-openhes-models-acquisition-listofvariable.md)

## GetVariables

```proto
GetVariables(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Variable)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.Variable`](model-io-clbs-openhes-models-acquisition-variable.md)

## UpdateVariable

```proto
UpdateVariable(io.clbs.openhes.models.acquisition.Variable)
```

- Input: [`io.clbs.openhes.models.acquisition.Variable`](model-io-clbs-openhes-models-acquisition-variable.md)

## DeleteVariable

```proto
DeleteVariable(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## AddRegisterToVariable

```proto
AddRegisterToVariable(io.clbs.openhes.models.acquisition.AddRegisterToVariableRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.AddRegisterToVariableRequest`](model-io-clbs-openhes-models-acquisition-addregistertovariablerequest.md)

## RemoveRegisterFromVariable

```proto
RemoveRegisterFromVariable(io.clbs.openhes.models.acquisition.RemoveRegisterFromVariableRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.RemoveRegisterFromVariableRequest`](model-io-clbs-openhes-models-acquisition-removeregisterfromvariablerequest.md)

