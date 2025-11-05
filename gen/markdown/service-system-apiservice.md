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

```proto
SetScreenConfig(io.clbs.openhes.models.system.SetScreenConfigRequest)
```

- Input: [`io.clbs.openhes.models.system.SetScreenConfigRequest`](model-io-clbs-openhes-models-system-setscreenconfigrequest.md)

## GetScreenConfig

```proto
GetScreenConfig(io.clbs.openhes.models.system.ScreenConfigSelector) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.system.ScreenConfigSelector`](model-io-clbs-openhes-models-system-screenconfigselector.md)
- Output: `google.protobuf.StringValue`

## SetScreenConfigs

```proto
SetScreenConfigs(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## GetScreenConfigs

```proto
GetScreenConfigs() returns (google.protobuf.StringValue)
```

- Output: `google.protobuf.StringValue`

## DeleteScreenConfig

```proto
DeleteScreenConfig(io.clbs.openhes.models.system.ScreenConfigSelector)
```

- Input: [`io.clbs.openhes.models.system.ScreenConfigSelector`](model-io-clbs-openhes-models-system-screenconfigselector.md)

