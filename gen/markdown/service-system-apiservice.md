# ApiService - System

## GetOpenIdConfiguration

The method returns the OIDC configuration, proxied directly from the configured OIDC service.
 All the authenticated endpoints shall be protected by token from this OIDC service.

```proto
GetOpenIdConfiguration() returns (io.clbs.openhes.models.system.OpenIdConfiguration)
```

- Output: [`io.clbs.openhes.models.system.OpenIdConfiguration`](model-io-clbs-openhes-models-system-openidconfiguration.md)

