# Model: io.clbs.openhes.models.crypto.CryptoSecret

Defines a specification of crypto secrets.

## Fields

| Field | Information |
| --- | --- |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>The access level for the secret (for example, `admin`, `user`). |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique secret key identifier. |
| createdAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the secret was created. |
| updatedAt | <b>Type:</b> `google.protobuf.Timestamp`<br><b>Description:</b><br>The timestamp when the secret was last updated. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The secret data. |

