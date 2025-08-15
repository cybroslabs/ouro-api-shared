# ApiService - System

## GetOpenIdConfiguration

Retrieves the details of the OpenId configuration, proxied directly from the configured OIDC service.
 All the authenticated endpoints shall be protected using a token issued by this OIDC service.

```proto
GetOpenIdConfiguration() returns (io.clbs.openhes.models.system.OpenIdConfiguration)
```

- Output: [`io.clbs.openhes.models.system.OpenIdConfiguration`](model-io-clbs-openhes-models-system-openidconfiguration.md)

