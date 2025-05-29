# ApiService - Configuration

## GetApplicationConfig

Gets the application configuration.

```proto
GetApplicationConfig() returns (io.clbs.openhes.models.system.ApplicationConfig)
```

- Output: [`io.clbs.openhes.models.system.ApplicationConfig`](model-io-clbs-openhes-models-system-applicationconfig.md)

## UpdateApplicationConfig

Updates the application configuration. The missing fields in the request will be kept unchanged.

```proto
UpdateApplicationConfig(io.clbs.openhes.models.system.ApplicationConfig)
```

- Input: [`io.clbs.openhes.models.system.ApplicationConfig`](model-io-clbs-openhes-models-system-applicationconfig.md)

## SynchronizeComponentConfig

Synchronizes the application configuration. The input value shall contain all the default values and also all known keys (with null values).
 The output value will contain currently set values inlcuding detauls which are not set.
 The missing values in the defaults will be deleted if has been set previously in the application configuration.

```proto
SynchronizeComponentConfig(io.clbs.openhes.models.system.ComponentConfigDescriptor) returns (io.clbs.openhes.models.system.ComponentConfig)
```

- Input: [`io.clbs.openhes.models.system.ComponentConfigDescriptor`](model-io-clbs-openhes-models-system-componentconfigdescriptor.md)
- Output: [`io.clbs.openhes.models.system.ComponentConfig`](model-io-clbs-openhes-models-system-componentconfig.md)

