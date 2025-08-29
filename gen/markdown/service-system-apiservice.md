# ApiService - System

## GetOpenIdConfiguration

Retrieves the details of the OpenId configuration, proxied directly from the configured OIDC service.
All the authenticated endpoints shall be protected using a token issued by this OIDC service.

```proto
GetOpenIdConfiguration() returns (io.clbs.openhes.models.system.OpenIdConfiguration)
```

- Output: [`io.clbs.openhes.models.system.OpenIdConfiguration`](model-io-clbs-openhes-models-system-openidconfiguration.md)

## GetTranslations

Retrieves the translation data.

```proto
GetTranslations(io.clbs.openhes.models.localization.GetTranslationsRequest) returns (io.clbs.openhes.models.localization.GetTranslationsResponse)
```

- Input: [`io.clbs.openhes.models.localization.GetTranslationsRequest`](model-io-clbs-openhes-models-localization-gettranslationsrequest.md)
- Output: [`io.clbs.openhes.models.localization.GetTranslationsResponse`](model-io-clbs-openhes-models-localization-gettranslationsresponse.md)

