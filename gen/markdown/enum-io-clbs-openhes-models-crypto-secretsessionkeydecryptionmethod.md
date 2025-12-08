# Enum: io.clbs.openhes.models.crypto.SecretSessionKeyDecryptionMethod

Defines the supported decryption methods for session keys used in key wrapping scenarios.
Session keys provide an additional layer of security for protecting cryptographic secrets.

## Options

| Value | Description |
| --- | --- |
| SECRET_SESSION_UNSPECIFIED | No session key is used. The secret data is encrypted directly with the master key. |
| SECRET_SESSION_PLAIN | Session key is stored in plain text (used when the session key itself is not sensitive). |
| SECRET_SESSION_RSA_OAEPM_GF1P | Session key is encrypted using RSA-OAEP with MGF1 and SHA-1. Used for secure key transport. |
