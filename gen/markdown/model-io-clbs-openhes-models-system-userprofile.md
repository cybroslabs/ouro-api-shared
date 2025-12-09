# Model: io.clbs.openhes.models.system.UserProfile

Defines the user profile containing identity and preferences for authenticated users.
 User profiles store both read-only attributes managed by the authentication system and user-controlled preferences.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique UUID of the user assigned by the authentication system (read-only). |
| displayName | <b>Type:</b> `string`<br><b>Description:</b><br>The user's full name or preferred display name shown in the UI.<br><b>Example:</b> "Jane Smith" |
| roles | <b>Type:</b> `string`<br><b>Description:</b><br>The authorization roles assigned to this user. Read-only, managed by the authentication system.<br><b>Example:</b> ["viewer"] |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The user's preferred UI language using IETF language tags.<br><b>Example:</b> "de-DE" |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The user's timezone for displaying dates and times using IANA timezone names.<br><b>Example:</b> "Asia/Tokyo" |

