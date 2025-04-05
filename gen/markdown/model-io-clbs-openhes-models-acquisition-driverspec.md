# Model: io.clbs.openhes.models.acquisition.DriverSpec

## Fields

| Field | Information |
| --- | --- |
| version | <b>Type:</b> `string`<br><b>Description:</b><br>The driver version. The format is not defined. The driver itself is versioned by the docker image tags so this value shall be either the same (set during the image build) or any useful user-readable version string. |
| listeningPort | <b>Type:</b> `uint32`<br><b>Description:</b><br>The port the driver's gRPC will listen on. |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The technical/internal ID of the driver. |
| maxConcurrentJobs | <b>Type:</b> `int32`<br><b>Description:</b><br>The maximum number of concurrent jobs the driver can handle. The value 0 is not allowed, the maximum number respect typical_mem_usage not to overgrow the memory resources! |
| maxCascadeDepth | <b>Type:</b> `uint32`<br><b>Description:</b><br>The maximum cascade depth the driver can handle. Number 1 means that the driver cannot handle cascading jobs, 2 means that the driver can handle cascading jobs with one level of depth, etc.<br> The value 0 means that the driver can handle any number of cascading jobs. |
| typicalMemUsage | <b>Type:</b> `int32`<br><b>Description:</b><br>The typical memory usage of the driver in MB. |
| templates | <b>Type:</b> [`io.clbs.openhes.models.acquisition.DriverTemplates`](model-io-clbs-openhes-models-acquisition-drivertemplates.md)<br><b>Description:</b><br>The connection and action templates. |
| displayName | <b>Type:</b> `string`<br><b>Description:</b><br>The display name of the driver. Must be in format '<manufacturer> <device_type> [<device_type_version>]'.<br> It must respect upper/lower characters.<br> The generic drivers, such as 'cybros labs generic', must be named as '<driver_company_name> generic'.<br><br> Examples: 'Addax NP73E', 'cybros labs generic', 'Landis+Gyr S650 v2' |

