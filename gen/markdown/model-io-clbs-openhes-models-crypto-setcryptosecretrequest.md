# Model: io.clbs.openhes.models.crypto.SetCryptoSecretRequest

## Fields

| Field | Information |
| --- | --- |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>Defines the resource type for which the field descriptor is defined, e.g., BULK, DEVICE, etc. |
| cryptoId | <b>Type:</b> `string`<br><b>Description:</b><br>The crypto ID of the secret to retrieve. |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>Access level for the secret, e.g., "admin", "user", etc. |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>Unique identifier for the secret key. |
| dataDecryptionSecretId | <b>Type:</b> `string`<br><b>Description:</b><br>Identifier for the key used to decrypt the secret data. May be unset if the data is not encrypted and vice versa. |
| dataDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretDataDesryptionMethod`](enum-io-clbs-openhes-models-crypto-secretdatadesryptionmethod.md)<br><b>Description:</b><br>Method used to decrypt the secret data. May be unset if the data is not encrypted and vice versa. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The secret data. |

