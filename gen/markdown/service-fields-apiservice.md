# ApiService - Fields

## CreateFieldDescriptor

The method to create a new field descriptor user-defined field descriptor.

```proto
CreateFieldDescriptor(io.clbs.openhes.models.common.CreateFieldDescriptorRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.common.CreateFieldDescriptorRequest`](model-io-clbs-openhes-models-common-createfielddescriptorrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateFieldDescriptor

The method to update the field descriptor.

```proto
UpdateFieldDescriptor(io.clbs.openhes.models.common.FieldDescriptor)
```

- Input: [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)

## DeleteFieldDescriptor

The method to delete the field descriptor.

```proto
DeleteFieldDescriptor(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## ListFieldDescriptors

The method to get the list of fields.

```proto
ListFieldDescriptors() returns (io.clbs.openhes.models.common.ListOfFieldDescriptor)
```

- Output: [`io.clbs.openhes.models.common.ListOfFieldDescriptor`](model-io-clbs-openhes-models-common-listoffielddescriptor.md)

