# ApiService - System

## GetLicenseRequestCode

The method returns the license request code if the license is not set. Otherwise it returns empty string.

```proto
GetLicenseRequestCode() returns (google.protobuf.StringValue)
```

- Output: `google.protobuf.StringValue`

## SetLicense

The method stored a new license key. Used only and only for air-gapped installations.

```proto
SetLicense(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## SetScreenConfig

Sets the screen configuration.

```proto
SetScreenConfig(io.clbs.openhes.models.system.SetScreenConfigRequest)
```

- Input: [`io.clbs.openhes.models.system.SetScreenConfigRequest`](model-io-clbs-openhes-models-system-setscreenconfigrequest.md)

## GetScreenConfig

Gets the screen configuration.

```proto
GetScreenConfig(io.clbs.openhes.models.system.ScreenConfigSelector) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.system.ScreenConfigSelector`](model-io-clbs-openhes-models-system-screenconfigselector.md)
- Output: `google.protobuf.StringValue`

## SetScreenConfigs

Sets multiple screen configurations at once, replacing any existing configurations.

```proto
SetScreenConfigs(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## GetScreenConfigs

Gets all screen configurations at once.

```proto
GetScreenConfigs() returns (google.protobuf.StringValue)
```

- Output: `google.protobuf.StringValue`

## DeleteScreenConfig

Deletes the specified screen configuration.

```proto
DeleteScreenConfig(io.clbs.openhes.models.system.ScreenConfigSelector)
```

- Input: [`io.clbs.openhes.models.system.ScreenConfigSelector`](model-io-clbs-openhes-models-system-screenconfigselector.md)

## GetObjectFlags

Retrieves the flags associated with the specified object.

```proto
GetObjectFlags(io.clbs.openhes.models.system.ObjectFlagsRequest) returns (io.clbs.openhes.models.system.ObjectFlagsResponse)
```

- Input: [`io.clbs.openhes.models.system.ObjectFlagsRequest`](model-io-clbs-openhes-models-system-objectflagsrequest.md)
- Output: [`io.clbs.openhes.models.system.ObjectFlagsResponse`](model-io-clbs-openhes-models-system-objectflagsresponse.md)

