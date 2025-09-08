# Enum: io.clbs.openhes.models.crypto.SecretSessionKeyDecryptionMethod

Defines the supported decryption methods for session keys.

## Options

| Value | Description |
| --- | --- |
| SECRET_SESSION_UNSPECIFIED | No session key decryption method. Used when the session key is not set or not used. |
| SECRET_SESSION_PLAIN | Plain session key. No session-key decryption required. |
| SECRET_SESSION_RSA_OAEPM_GF1P | RSA-OAEP-MGF1P decryption method. Used for session keys. |
