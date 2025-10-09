# Model: io.clbs.openhes.models.localization.UpdateTranslationsRequest

Defines the request for updating translations.

## Fields

| Field | Information |
| --- | --- |
| version | <b>Type:</b> `string`<br><b>Description:</b><br>The version of the translations, corresponding to the application or plugin version. |
| application | <b>Type:</b> `string`<br><b>Description:</b><br>The application or plugin identifier where the translations are to be updated. Can be empty for core application. |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The IETF language tag (for example, `en-US`, `fr-FR`) for which translations are to be updated. |
| translations | <b>Type:</b> map<`string`, `string`><br><b>Description:</b><br>A map of keys to their corresponding translated strings. |

