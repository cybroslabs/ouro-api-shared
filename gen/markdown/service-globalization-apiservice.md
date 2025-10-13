# ApiService - Globalization

## GetTranslations

Retrieves the translation data.

```proto
GetTranslations(io.clbs.openhes.models.localization.GetTranslationsRequest) returns (io.clbs.openhes.models.localization.GetTranslationsResponse)
```

- Input: [`io.clbs.openhes.models.localization.GetTranslationsRequest`](model-io-clbs-openhes-models-localization-gettranslationsrequest.md)
- Output: [`io.clbs.openhes.models.localization.GetTranslationsResponse`](model-io-clbs-openhes-models-localization-gettranslationsresponse.md)

## SetTranslationMissing

Indicates that a translation is missing for the specified language and key.

```proto
SetTranslationMissing(io.clbs.openhes.models.localization.MissingTranslationRequest)
```

- Input: [`io.clbs.openhes.models.localization.MissingTranslationRequest`](model-io-clbs-openhes-models-localization-missingtranslationrequest.md)

## UpdateTranslations

Updates the translations for a specific language. Existing translations for the specified language will be replaced with the new ones provided in the request.

```proto
UpdateTranslations(io.clbs.openhes.models.localization.UpdateTranslationsRequest)
```

- Input: [`io.clbs.openhes.models.localization.UpdateTranslationsRequest`](model-io-clbs-openhes-models-localization-updatetranslationsrequest.md)

