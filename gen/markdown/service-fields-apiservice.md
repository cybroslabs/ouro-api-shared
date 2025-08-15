# ApiService - Fields

## CreateFieldDescriptor

Creates a new field descriptor. Returns the identifier of the newly created field descriptor.

```proto
CreateFieldDescriptor(io.clbs.openhes.models.common.CreateFieldDescriptorRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.common.CreateFieldDescriptorRequest`](model-io-clbs-openhes-models-common-createfielddescriptorrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateFieldDescriptor

Updates the details of an existing field descriptor. Fields that are omitted from the request will be left unchanged.

```proto
UpdateFieldDescriptor(io.clbs.openhes.models.common.FieldDescriptor)
```

- Input: [`io.clbs.openhes.models.common.FieldDescriptor`](model-io-clbs-openhes-models-common-fielddescriptor.md)

## DeleteFieldDescriptor

Deletes the specified field descriptor.

```proto
DeleteFieldDescriptor(io.clbs.openhes.models.common.FieldDescriptorSelector)
```

- Input: [`io.clbs.openhes.models.common.FieldDescriptorSelector`](model-io-clbs-openhes-models-common-fielddescriptorselector.md)

## ListFieldDescriptors

Retrieves a paginated list of field descriptors based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListFieldDescriptors() returns (io.clbs.openhes.models.common.ListOfFieldDescriptor)
```

- Output: [`io.clbs.openhes.models.common.ListOfFieldDescriptor`](model-io-clbs-openhes-models-common-listoffielddescriptor.md)

## ListFieldDescriptorOptions

Retrieves a paginated list of available options for the field descriptor based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListFieldDescriptorOptions(io.clbs.openhes.models.common.ListFieldDescriptorOptionsRequest) returns (io.clbs.openhes.models.common.FieldDescriptorOptions)
```

- Input: [`io.clbs.openhes.models.common.ListFieldDescriptorOptionsRequest`](model-io-clbs-openhes-models-common-listfielddescriptoroptionsrequest.md)
- Output: [`io.clbs.openhes.models.common.FieldDescriptorOptions`](model-io-clbs-openhes-models-common-fielddescriptoroptions.md)

