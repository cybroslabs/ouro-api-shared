# Model: io.clbs.openhes.models.crypto.GetCryptoSecretRequest

Defines a specification for retrieving crypto secrets.

## Fields

| Field | Information |
| --- | --- |
| objectType | <b>Type:</b> [`io.clbs.openhes.models.common.ObjectType`](enum-io-clbs-openhes-models-common-objecttype.md)<br><b>Description:</b><br>Thehe resource type for which the secret is defined (for example, `BULK`, `DEVICE`). |
| driverType | <b>Type:</b> `string`<br><b>Description:</b><br>The driver type for which the secret is requested (for example, `METERCONTROL_ST402D_DLMS`). |
| cryptoId | <b>Type:</b> `string`<br><b>Description:</b><br>The crypto identifier of the secret to retrieve. |

