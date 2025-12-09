# Model: io.clbs.openhes.models.crypto.GetCryptoSecretRequest

Defines a specification for retrieving crypto secrets.

## Fields

| Field | Information |
| --- | --- |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>The resource type for which the secret is defined.<br><b>Example:</b> OBJECT_TYPE_DEVICE |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type for which the secret is requested.<br><b>Example:</b> "LANDISGYR_E650_DLMS_SN" |
| cryptoId | <b>Type:</b> `string - UUID`<br><b>Description:</b><br>The crypto identifier of the secret to retrieve. |

