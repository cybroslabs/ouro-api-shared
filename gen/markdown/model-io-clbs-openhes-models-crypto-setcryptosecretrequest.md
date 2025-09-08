# Model: io.clbs.openhes.models.crypto.SetCryptoSecretRequest

Defines a specification for creating or updating crypto secrets.

## Fields

| Field | Information |
| --- | --- |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>Thehe resource type for which the secret is defined (for example, `BULK`, `DEVICE`). |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type for which the secret is requested (for example, `METERCONTROL_ST402D_DLMS`). |
| cryptoId | <b>Type:</b> `string`<br><b>Description:</b><br>The crypto identifier of the secret to retrieve. |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>The access level for the secret (for example, `admin`, `user`). |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique secret key identifier. |
| decryptionSecretId | <b>Type:</b> `string`<br><b>Description:</b><br>The uniqie identifier for the key used to decrypt either the session key or the data directly. |
| sessionKeyDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretSessionKeyDecryptionMethod`](enum-io-clbs-openhes-models-crypto-secretsessionkeydecryptionmethod.md)<br><b>Description:</b><br>the method used to decrypt the session key. May be unset if the session key is not set or not used. |
| sessionKey | <b>Type:</b> `bytes`<br><b>Description:</b><br>The session key (usually encrypted) used to decrypt the data. |
| dataDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretDataDecryptionMethod`](enum-io-clbs-openhes-models-crypto-secretdatadecryptionmethod.md)<br><b>Description:</b><br>The method used to decrypt the secret data. May be unset if the data is not encrypted. |
| dataDecryptionIv | <b>Type:</b> `bytes`<br><b>Description:</b><br>The initialization vector for the decryption method, if applicable. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The secret data. |

