# Model: io.clbs.openhes.models.crypto.CryptoSecret

Defines a cryptographic secret containing sensitive credentials for device communication.
 Secrets include authentication keys, encryption keys, and passwords used to establish secure connections with devices.

## Fields

| Field | Information |
| --- | --- |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>The security access level this secret provides (e.g., "admin", "user", "read-only"). Maps to DLMS/COSEM association levels. |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier for this secret key. Used to distinguish between multiple keys for the same device. |
| createdAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when this secret was initially created in the system. |
| updatedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when this secret was last modified or rotated. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The actual secret data (e.g., encryption key, password). This field contains the decrypted value ready for use. |

