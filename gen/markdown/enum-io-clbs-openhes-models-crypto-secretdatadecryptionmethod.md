# Enum: io.clbs.openhes.models.crypto.SecretDataDecryptionMethod

Defines the supported decryption methods for cryptographic secret data.
These methods are used to decrypt device authentication keys and other sensitive credentials.

## Options

| Value | Description |
| --- | --- |
| SECRET_DATA_UNSPECIFIED | Unspecified decryption method (invalid, should not be used). |
| SECRET_DATA_PLAIN | Data is stored in plain text without encryption (not recommended for production). |
| SECRET_DATA_AES256CBC | Data is encrypted using AES-256 in CBC mode. Requires an initialization vector. |
