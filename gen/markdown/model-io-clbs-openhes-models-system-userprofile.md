# Model: io.clbs.openhes.models.system.UserProfile

Defines the user profile containing identity and preferences for authenticated users.
 User profiles store both read-only attributes managed by the authentication system and user-controlled preferences.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The unique UUID of the user assigned by the authentication system (read-only). |
| displayName | <b>Type:</b> `string`<br><b>Description:</b><br>The user's full name or preferred display name shown in the UI. |
| roles | <b>Type:</b> `string`<br><b>Description:</b><br>The authorization roles assigned to this user (e.g., "admin", "operator"). Read-only, managed by the authentication system. |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The user's preferred UI language using IETF language tags (e.g., "en-US", "cs-CZ"). |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The user's timezone for displaying dates and times using IANA timezone names (e.g., "America/New_York", "Europe/Prague"). |

