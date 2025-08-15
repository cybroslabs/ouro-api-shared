# ApiService - Cron Jobs

## CreateCronJob

Creates a new cron job. Returns the identifier of the newly created cron job.

```proto
CreateCronJob(io.clbs.openhes.models.cronjobs.CreateCronJobRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.cronjobs.CreateCronJobRequest`](model-io-clbs-openhes-models-cronjobs-createcronjobrequest.md)
- Output: `google.protobuf.StringValue`

## ListCronJobs

Retrieves a paginated list of cron jobs based on the specified criteria. The page size and page number (zero-based) are defined in the request.

```proto
ListCronJobs(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.cronjobs.ListOfCronJob)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.cronjobs.ListOfCronJob`](model-io-clbs-openhes-models-cronjobs-listofcronjob.md)

## GetCronJob

Retrieves the details of the specified cron job.

```proto
GetCronJob(google.protobuf.StringValue) returns (io.clbs.openhes.models.cronjobs.CronJob)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.cronjobs.CronJob`](model-io-clbs-openhes-models-cronjobs-cronjob.md)

## UpdateCronJob

Updates the details of an existing cron job.

```proto
UpdateCronJob(io.clbs.openhes.models.cronjobs.CronJob)
```

- Input: [`io.clbs.openhes.models.cronjobs.CronJob`](model-io-clbs-openhes-models-cronjobs-cronjob.md)

## DeleteCronJob

Deletes the specified cron job.

```proto
DeleteCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## RunCronJob

Runs the specified cron job immediately.

```proto
RunCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## PauseCronJob

Pauses the specified cron job.

```proto
PauseCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## ResumeCronJob

Resumes a paused cron job and restores its scheduled execution.

```proto
ResumeCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

