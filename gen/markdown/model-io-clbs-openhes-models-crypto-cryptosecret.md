# Model: io.clbs.openhes.models.crypto.CryptoSecret

Defines a cryptographic secret containing sensitive credentials for device communication.
 Secrets include authentication keys, encryption keys, and passwords used to establish secure connections with devices.

## Fields

| Field | Information |
| --- | --- |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>The security access level this secret relates to. Maps to DLMS/COSEM association levels or similar concepts in other protocols.<br><b>Example:</b> "cert" |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier for this secret value. Used to distinguish between multiple crypto secrets within the same device.<br><b>Example:</b> "CERT" |
| createdAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when this secret was initially created in the system. |
| updatedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when this secret was last modified or rotated. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The actual secret data (e.g., encryption key, password). This field contains the decrypted value ready for use. |

