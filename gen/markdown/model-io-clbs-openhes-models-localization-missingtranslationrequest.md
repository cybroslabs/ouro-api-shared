# Model: io.clbs.openhes.models.localization.MissingTranslationRequest

Indicates that a translation is missing for the specified language and key.

## Fields

| Field | Information |
| --- | --- |
| applicationId | <b>Type:</b> `string`<br><b>Description:</b><br>The application or ui-plugin identifier where the translation is missing. Can be empty for core application. |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The IETF language tag (for example, `en-US`, `fr-FR`) for which a translation is missing. |
| key | <b>Type:</b> `string`<br><b>Description:</b><br>The key for which the translation is missing. |

