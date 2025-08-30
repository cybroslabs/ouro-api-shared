# ApiService - General

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

