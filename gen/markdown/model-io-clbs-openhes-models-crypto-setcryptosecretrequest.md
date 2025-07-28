# Model: io.clbs.openhes.models.crypto.SetCryptoSecretRequest

## Fields

| Field | Information |
| --- | --- |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>Defines the resource type for which the field descriptor is defined, e.g., BULK, DEVICE, etc. |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type for which the secret is requested, e.g., "METERCONTROL_ST402D_DLMS" |
| cryptoId | <b>Type:</b> `string`<br><b>Description:</b><br>The crypto ID of the secret to retrieve. |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>Access level for the secret, e.g., "admin", "user", etc. |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>Unique identifier for the secret key. |
| decryptionSecretId | <b>Type:</b> `string`<br><b>Description:</b><br>Identifier for the key used to decrypt the either the session key or the data directly. |
| sessionKeyDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretSessionKeyDecryptionMethod`](enum-io-clbs-openhes-models-crypto-secretsessionkeydecryptionmethod.md)<br><b>Description:</b><br>Method used to decrypt the session key. May be unset if the session key is unset and no used. |
| sessionKey | <b>Type:</b> `bytes`<br><b>Description:</b><br>The session key, usually encrypted, used to decrypt the data. |
| dataDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretDataDecryptionMethod`](enum-io-clbs-openhes-models-crypto-secretdatadecryptionmethod.md)<br><b>Description:</b><br>Method used to decrypt the secret data. May be unset if the data is not encrypted and vice versa. |
| dataDecryptionIv | <b>Type:</b> `bytes`<br><b>Description:</b><br>Initialization vector for the decryption method, if applicable. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The secret data. |

