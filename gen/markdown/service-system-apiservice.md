# ApiService - System

## GetLicenseRequestCode

Retrieves the license request code for the system. This code is used to generate a license for air-gapped installations. Returns an empty string if a valid license is already installed.

```proto
GetLicenseRequestCode() returns (google.protobuf.StringValue)
```

- Output: `google.protobuf.StringValue`

## SetLicense

Installs a new license key. This method is intended for air-gapped installations where the license cannot be activated through online channels. The license key must be obtained from Cybroslabs using the license request code.

```proto
SetLicense(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## SetScreenConfig

Sets or updates a screen configuration. Screen configurations store UI-specific settings and layouts for different screens or views in the application.

```proto
SetScreenConfig(io.clbs.openhes.models.system.SetScreenConfigRequest)
```

- Input: [`io.clbs.openhes.models.system.SetScreenConfigRequest`](model-io-clbs-openhes-models-system-setscreenconfigrequest.md)

## GetScreenConfig

Retrieves a specific screen configuration based on the selector. Returns the configuration as a JSON string.

```proto
GetScreenConfig(io.clbs.openhes.models.system.ScreenConfigSelector) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.system.ScreenConfigSelector`](model-io-clbs-openhes-models-system-screenconfigselector.md)
- Output: `google.protobuf.StringValue`

## SetScreenConfigs

Sets multiple screen configurations at once, replacing all existing configurations. This is useful for bulk updates or restoring configurations from backup.

```proto
SetScreenConfigs(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## GetScreenConfigs

Retrieves all screen configurations at once as a JSON string. This is useful for exporting or backing up configurations.

```proto
GetScreenConfigs() returns (google.protobuf.StringValue)
```

- Output: `google.protobuf.StringValue`

## DeleteScreenConfig

Deletes a specific screen configuration identified by the selector.

```proto
DeleteScreenConfig(io.clbs.openhes.models.system.ScreenConfigSelector)
```

- Input: [`io.clbs.openhes.models.system.ScreenConfigSelector`](model-io-clbs-openhes-models-system-screenconfigselector.md)

## GetObjectFlags

Retrieves feature flags and capability indicators associated with the specified object. These flags control what operations or features are available for the object.

```proto
GetObjectFlags(io.clbs.openhes.models.system.ObjectFlagsRequest) returns (io.clbs.openhes.models.system.ObjectFlagsResponse)
```

- Input: [`io.clbs.openhes.models.system.ObjectFlagsRequest`](model-io-clbs-openhes-models-system-objectflagsrequest.md)
- Output: [`io.clbs.openhes.models.system.ObjectFlagsResponse`](model-io-clbs-openhes-models-system-objectflagsresponse.md)

## GetSbom

Retrieves the software bill of materials (SBOM) information in CycloneDX JSON format.

```proto
GetSbom() returns (google.protobuf.StringValue)
```

- Output: `google.protobuf.StringValue`

