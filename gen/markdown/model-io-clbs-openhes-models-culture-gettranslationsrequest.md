# Model: io.clbs.openhes.models.culture.GetTranslationsRequest

Defines the request message for retrieving translations based on language (and keys).

## Fields

| Field | Information |
| --- | --- |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The IETF language tag (e.g., "en-US", "fr-FR") for which translations are requested. |
| keys | <b>Type:</b> `string`<br><b>Description:</b><br>List of keys for which translations are requested. If empty, all translations for the specified language will be returned. |

