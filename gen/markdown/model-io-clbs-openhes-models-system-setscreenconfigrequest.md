# Model: io.clbs.openhes.models.system.SetScreenConfigRequest

Defines a request for setting or updating UI screen configuration.
 Screen configurations store user-specific or system-wide UI layout and settings for different application screens.

## Fields

| Field | Information |
| --- | --- |
| applicationId | <b>Type:</b> `string`<br><b>Description:</b><br>The application or UI plugin identifier.<br><b>Example:</b> "scada-plugin" |
| screenId | <b>Type:</b> `string`<br><b>Description:</b><br>The screen identifier within the application.<br><b>Example:</b> "bulk-overview" |
| json | <b>Type:</b> `string`<br><b>Description:</b><br>Configuration as a JSON-encoded string. Useful for direct storage from frontend. |
| raw | <b>Type:</b> `google.protobuf.Struct`<br><b>Description:</b><br>Configuration as a structured protobuf Struct. Used for programmatic access. |

