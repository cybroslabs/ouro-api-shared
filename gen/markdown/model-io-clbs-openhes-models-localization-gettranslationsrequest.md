# Model: io.clbs.openhes.models.localization.GetTranslationsRequest

Defines the request for retrieving translations based on language and keys.

## Fields

| Field | Information |
| --- | --- |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The IETF language tag (for example, `en-US`, `fr-FR`) for which translations are requested. |
| keys | <b>Type:</b> `string`<br><b>Description:</b><br>The list of keys for which translations are requested. If empty, all translations for the specified language are returned. |

