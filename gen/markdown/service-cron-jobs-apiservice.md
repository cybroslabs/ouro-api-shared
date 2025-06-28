# ApiService - Cron Jobs

## CreateCronJob

The method to create a new cron job.

```proto
CreateCronJob(io.clbs.openhes.models.cronjobs.CreateCronJobRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.cronjobs.CreateCronJobRequest`](model-io-clbs-openhes-models-cronjobs-createcronjobrequest.md)
- Output: `google.protobuf.StringValue`

## ListCronJobs

The method to get the list of cron jobs.

```proto
ListCronJobs(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.cronjobs.ListOfCronJob)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.cronjobs.ListOfCronJob`](model-io-clbs-openhes-models-cronjobs-listofcronjob.md)

## GetCronJob

The method to get the cron job.

```proto
GetCronJob(google.protobuf.StringValue) returns (io.clbs.openhes.models.cronjobs.CronJob)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.cronjobs.CronJob`](model-io-clbs-openhes-models-cronjobs-cronjob.md)

## UpdateCronJob

The method to update the cron job.

```proto
UpdateCronJob(io.clbs.openhes.models.cronjobs.CronJob)
```

- Input: [`io.clbs.openhes.models.cronjobs.CronJob`](model-io-clbs-openhes-models-cronjobs-cronjob.md)

## DeleteCronJob

The method to delete the cron job.

```proto
DeleteCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## RunCronJob

The method to run the cron job immediately.

```proto
RunCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## PauseCronJob

The method to pause the cron job.

```proto
PauseCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## ResumeCronJob

The method to resume the cron job.

```proto
ResumeCronJob(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

