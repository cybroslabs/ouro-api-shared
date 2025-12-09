# Model: io.clbs.openhes.models.crypto.SetCryptoSecretRequest

Defines a specification for creating or updating crypto secrets.

## Fields

| Field | Information |
| --- | --- |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>The resource type for which the secret is defined.<br><b>Example:</b> OBJECT_TYPE_DEVICE |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type for which the secret is created or updated.<br><b>Example:</b> "LANDISGYR_E650_DLMS_SN" |
| cryptoId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The crypto identifier of the secret to create or update. |
| accessLevel | <b>Type:</b> `string`<br><b>Description:</b><br>The security access level for this secret value. Maps to DLMS/COSEM association levels or similar concepts in other protocols.<br><b>Example:</b> "cert" |
| keyId | <b>Type:</b> `string`<br><b>Description:</b><br>The unique identifier for this secret value. Used to distinguish between multiple crypto secrets within the same device.<br><b>Example:</b> "CERT" |
| decryptionSecretId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The unique identifier for the key used to decrypt either the session key or the data directly. |
| sessionKeyDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretSessionKeyDecryptionMethod`](enum-io-clbs-openhes-models-crypto-secretsessionkeydecryptionmethod.md)<br><b>Description:</b><br>The method used to decrypt the session key. May be unset if the session key is not set or not used. |
| sessionKey | <b>Type:</b> `bytes`<br><b>Description:</b><br>The session key (usually encrypted) used to decrypt the data. |
| dataDecryptionMethod | <b>Type:</b> [`io.clbs.openhes.models.crypto.SecretDataDecryptionMethod`](enum-io-clbs-openhes-models-crypto-secretdatadecryptionmethod.md)<br><b>Description:</b><br>The method used to decrypt the secret data. May be unset if the data is not encrypted. |
| dataDecryptionIv | <b>Type:</b> `bytes`<br><b>Description:</b><br>The initialization vector for the decryption method, if applicable. |
| data | <b>Type:</b> `bytes`<br><b>Description:</b><br>The secret data. |

