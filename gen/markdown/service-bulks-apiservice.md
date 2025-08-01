# ApiService - Bulks

## ListBulks

Retrieves the list of bulks. The list of bulks is paginated. The page size is defined in the request. The page number is 0-based.
 The list contains both the proxy bulks and the regular bulks.

```proto
ListBulks(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfBulk)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfBulk`](model-io-clbs-openhes-models-acquisition-listofbulk.md)

## ListBulkJobs

Retrieves the list of jobs. The list of jobs is paginated. The page size is defined in the request. The page number is 0-based.
 The listing can be used for both proxy bulks and regular bulks.

```proto
ListBulkJobs(io.clbs.openhes.models.acquisition.ListBulkJobsRequest) returns (io.clbs.openhes.models.acquisition.ListOfBulkJob)
```

- Input: [`io.clbs.openhes.models.acquisition.ListBulkJobsRequest`](model-io-clbs-openhes-models-acquisition-listbulkjobsrequest.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfBulkJob`](model-io-clbs-openhes-models-acquisition-listofbulkjob.md)

## GetBulkJob

Retrieves the job status. It can be used for jobs related to both proxy and regular bulks.

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

Cancels the bulk of jobs. It can be used for both proxy and regular bulks.

```proto
CancelBulk(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## CancelBulkJobs

Cancels the job(s) identified by the job identifier(s).

```proto
CancelBulkJobs(io.clbs.openhes.models.common.ListOfId)
```

- Input: [`io.clbs.openhes.models.common.ListOfId`](model-io-clbs-openhes-models-common-listofid.md)

## CreateProxyBulk

Starts a new proxy bulk. The proxy bolk is a collection of jobs where each job represents a single device. Devices must be fully defined in the request.

```proto
CreateProxyBulk(io.clbs.openhes.models.acquisition.CreateProxyBulkRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateProxyBulkRequest`](model-io-clbs-openhes-models-acquisition-createproxybulkrequest.md)
- Output: `google.protobuf.StringValue`

## GetProxyBulk

Retrieves the proxy bulk info and status.

```proto
GetProxyBulk(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ProxyBulk)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.ProxyBulk`](model-io-clbs-openhes-models-acquisition-proxybulk.md)

## CreateBulk

Starts a new bulk. The bulk is a collection of jobs where each jobs represents a single device. Devices that are part of the bulk are identified either as a list of registered device identifiers or as a group identifier.

```proto
CreateBulk(io.clbs.openhes.models.acquisition.CreateBulkRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateBulkRequest`](model-io-clbs-openhes-models-acquisition-createbulkrequest.md)
- Output: `google.protobuf.StringValue`

## GetBulk

Retrieves the bulk info and status.

```proto
GetBulk(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Bulk)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.Bulk`](model-io-clbs-openhes-models-acquisition-bulk.md)

## UpdateBulk

Updates the bulk metadata. The metadata is used to store additional information about the job.

```proto
UpdateBulk(io.clbs.openhes.models.common.UpdateMetadata)
```

- Input: [`io.clbs.openhes.models.common.UpdateMetadata`](model-io-clbs-openhes-models-common-updatemetadata.md)

## GetDeviceBulks

Retrieves the list of bulk jobs related to given device in the specified time range. All the parameters are required.

```proto
GetDeviceBulks(io.clbs.openhes.models.acquisition.GetDeviceBulksRequest) returns (io.clbs.openhes.models.acquisition.DeviceBulks)
```

- Input: [`io.clbs.openhes.models.acquisition.GetDeviceBulksRequest`](model-io-clbs-openhes-models-acquisition-getdevicebulksrequest.md)
- Output: [`io.clbs.openhes.models.acquisition.DeviceBulks`](model-io-clbs-openhes-models-acquisition-devicebulks.md)

