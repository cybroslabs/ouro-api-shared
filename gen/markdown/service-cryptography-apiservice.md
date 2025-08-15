# ApiService - Cryptography

## GetCryptoSecret

Retrieves a cryptographic secret based on the specified request parameters.

```proto
GetCryptoSecret(io.clbs.openhes.models.crypto.GetCryptoSecretRequest) returns (io.clbs.openhes.models.crypto.CryptoSecrets)
```

- Input: [`io.clbs.openhes.models.crypto.GetCryptoSecretRequest`](model-io-clbs-openhes-models-crypto-getcryptosecretrequest.md)
- Output: [`io.clbs.openhes.models.crypto.CryptoSecrets`](model-io-clbs-openhes-models-crypto-cryptosecrets.md)

## SetCryptoSecret

Creates a cryptographic the secret. If a secret with the same identifier already exists, it will be replaced.

```proto
SetCryptoSecret(io.clbs.openhes.models.crypto.SetCryptoSecretRequest)
```

- Input: [`io.clbs.openhes.models.crypto.SetCryptoSecretRequest`](model-io-clbs-openhes-models-crypto-setcryptosecretrequest.md)

