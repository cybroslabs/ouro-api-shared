# ApiService - Bulks

## ListBulks

Retrieves a paginated list of bulks based on the specified criteria. The page size and page number (zero-based) are defined in the request.
 The list contains both proxy and regular bulks.

```proto
ListBulks(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfBulk)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfBulk`](model-io-clbs-openhes-models-acquisition-listofbulk.md)

## ListBulkJobs

Retrieves a paginated list of jobs based on the specified criteria. The page size and page number (zero-based) are defined in the request.
 The listing can be used for jobs from both proxy and regular bulks.

```proto
ListBulkJobs(io.clbs.openhes.models.acquisition.ListBulkJobsRequest) returns (io.clbs.openhes.models.acquisition.ListOfBulkJob)
```

- Input: [`io.clbs.openhes.models.acquisition.ListBulkJobsRequest`](model-io-clbs-openhes-models-acquisition-listbulkjobsrequest.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfBulkJob`](model-io-clbs-openhes-models-acquisition-listofbulkjob.md)

## GetBulkJob

Retrieves the details of the specified job. It can be used for jobs from both proxy and regular bulks.

```proto
GetBulkJob(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.BulkJob)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.BulkJob`](model-io-clbs-openhes-models-acquisition-bulkjob.md)

## UpdateBulkJob

Updates the job metadata. The metadata is used to store additional information about the job.

```proto
UpdateBulkJob(io.clbs.openhes.models.common.UpdateMetadata)
```

- Input: [`io.clbs.openhes.models.common.UpdateMetadata`](model-io-clbs-openhes-models-common-updatemetadata.md)

## CancelBulk

Cancels the specified job bulk. It can be used for both proxy and regular bulks.

```proto
CancelBulk(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## CancelBulkJobs

Cancels the specified jobs in an existing bulk.

```proto
CancelBulkJobs(io.clbs.openhes.models.common.ListOfId)
```

- Input: [`io.clbs.openhes.models.common.ListOfId`](model-io-clbs-openhes-models-common-listofid.md)

## CreateProxyBulk

Creates a new proxy bulk. The proxy bulk is a collection of jobs where each job represents a single device. Devices must be fully defined in the request.

```proto
CreateProxyBulk(io.clbs.openhes.models.acquisition.CreateProxyBulkRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateProxyBulkRequest`](model-io-clbs-openhes-models-acquisition-createproxybulkrequest.md)
- Output: `google.protobuf.StringValue`

## GetProxyBulk

Retrieves the details of the specified proxy bulk.

```proto
GetProxyBulk(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ProxyBulk)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.ProxyBulk`](model-io-clbs-openhes-models-acquisition-proxybulk.md)

## CreateBulk

Creates a new bulk. The bulk is a collection of jobs where each job represents a single device. Devices that are part of the bulk are identified either as a list of registered device identifiers or as a group identifier.

```proto
CreateBulk(io.clbs.openhes.models.acquisition.CreateBulkRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateBulkRequest`](model-io-clbs-openhes-models-acquisition-createbulkrequest.md)
- Output: `google.protobuf.StringValue`

## GetBulk

Retrieves the the details of the specified bulk.

```proto
GetBulk(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Bulk)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.Bulk`](model-io-clbs-openhes-models-acquisition-bulk.md)

## UpdateBulk

Updates the metadata of an existing bulk. The metadata is used to store additional information about the job.

```proto
UpdateBulk(io.clbs.openhes.models.common.UpdateMetadata)
```

- Input: [`io.clbs.openhes.models.common.UpdateMetadata`](model-io-clbs-openhes-models-common-updatemetadata.md)

## GetDeviceBulks

Retrieves the list of bulk jobs related to a given device within the specified time range. All parameters are required.

```proto
GetDeviceBulks(io.clbs.openhes.models.acquisition.GetDeviceBulksRequest) returns (io.clbs.openhes.models.acquisition.DeviceBulks)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceBulksRequest`](model-io-clbs-openhes-models-acquisition-getdevicebulksrequest.md)
- Output: [`io.clbs.openhes.models.acquisition.DeviceBulks`](model-io-clbs-openhes-models-acquisition-devicebulks.md)

