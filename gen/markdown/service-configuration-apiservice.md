# ApiService - Configuration

## GetApplicationConfig

Retrieves the current application configuration settings.

```proto
GetApplicationConfig() returns (io.clbs.openhes.models.system.ApplicationConfigDescriptor)
```

- Output: [`io.clbs.openhes.models.system.ApplicationConfigDescriptor`](model-io-clbs-openhes-models-system-applicationconfigdescriptor.md)

## UpdateApplicationConfig

Updates the details of an existing application configuration. Fields that are omitted from the request will be left unchanged.

```proto
UpdateApplicationConfig(io.clbs.openhes.models.system.ApplicationConfig)
```

- Input: [`io.clbs.openhes.models.system.ApplicationConfig`](model-io-clbs-openhes-models-system-applicationconfig.md)

## SynchronizeComponentConfig

Synchronizes the application configuration. The input value shall contain all default values and all known keys (even with null values).
The output value will contain currently set values, including details that are not set.
Values missing from the defaults will be deleted if they were previously set in the application configuration.

```proto
SynchronizeComponentConfig(io.clbs.openhes.models.system.ComponentConfigDescriptor) returns (io.clbs.openhes.models.system.ComponentConfig)
```

- Input: [`io.clbs.openhes.models.system.ComponentConfigDescriptor`](model-io-clbs-openhes-models-system-componentconfigdescriptor.md)
- Output: [`io.clbs.openhes.models.system.ComponentConfig`](model-io-clbs-openhes-models-system-componentconfig.md)

