# Model: io.clbs.openhes.models.system.UserProfile

Defines the user information structure.

## Fields

| Field | Information |
| --- | --- |
| id | <b>Type:</b> `string`<br><b>Description:</b><br>The UUID of the user. |
| displayName | <b>Type:</b> `string`<br><b>Description:</b><br>The display name of the user. |
| roles | <b>Type:</b> `string`<br><b>Description:</b><br>The list of user roles. This is read-only and managed by the system. |
| ietfLanguageTag | <b>Type:</b> `string`<br><b>Description:</b><br>The preferred user interface language. |
| timezone | <b>Type:</b> `string`<br><b>Description:</b><br>The timezone associated with the device. For example, `"America/New_York"`, `"Europe/Prague"`, `"CET"`, `"GMT"`, `"Etc/GMT+2"`. |

