# Model: io.clbs.openhes.models.acquisition.DriverSpec

Defines the driver specification.

## Fields

| Field | Information |
| --- | --- |
| version | <b>Type:</b> `string`<br><b>Description:</b><br>The driver version. The format is not defined. Typically matches the docker image tag or another user-readable version string. |
| listeningPort | <b>Type:</b> `uint32`<br><b>Description:</b><br>The port on which the driver's gRPC listens. |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The technical/internal ID of the driver. |
| maxConcurrentJobs | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of concurrent jobs the driver can handle. A value of `0` is not allowed. The maximum value should respect `typical_mem_usage` to avoid exceeding memory resources! |
| maxCascadeDepth | <b>Type:</b> `uint32`<br><b>Description:</b><br>The maximum cascade depth the driver can handle. Value `1` means that cascading jobs are not supported. Value `2` means one level of cascading is allowed, and so on.<br>Value `0` means that the driver can handle any number of cascading jobs. |
| typicalMemUsage | <b>Type:</b> `int32`<br><b>Description:</b><br>The typical memory usage of the driver in `MB`. |
| templates | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DriverTemplates`](model-io-clbs-openhes-models-acquisition-drivertemplates.md)<br><b>Description:</b><br>The connection and action templates. |
| displayName | <b>Type:</b> `string`<br><b>Description:</b><br>The display name of the driver. Must be in the following format: `<manufacturer> <device_type> [<device_type_version>]`.<br>Must respect upper/lower case.<br>Generic drivers, such as `cybros labs generic`, must be in the following format: `<driver_company_name> generic`.<br><b>Example:</b> `Addax NP73E`, `cybros labs generic`, `Landis+Gyr S650 v2` |

