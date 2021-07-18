
# Table: gcp_kms_keyring_crypto_keys
A CryptoKey represents a logical key that can be used for cryptographic operations.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|keyring_cq_id|uuid|Unique ID of gcp_kms_keyrings table (FK)|
|project_id|text|GCP Project Id of the resource|
|location|text|Location of the resource|
|policy|jsonb|Access control policy for a resource|
|create_time|timestamp without time zone|The time at which this CryptoKey was created|
|labels|jsonb|Labels with user-defined metadata For more information, see Labeling Keys (https://cloudgooglecom/kms/docs/labeling-keys)|
|name|text|The resource name for this CryptoKey in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*`|
|next_rotation_time|timestamp without time zone|At next_rotation_time, the Key Management Service will automatically: 1 Create a new version of this CryptoKey 2 Mark the new version as primary Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time Keys with purpose ENCRYPT_DECRYPT support automatic rotation For other keys, this field must be omitted|
|primary_algorithm|text|The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports  Possible values:   "CRYPTO_KEY_VERSION_ALGORITHM_UNSPECIFIED" - Not specified   "GOOGLE_SYMMETRIC_ENCRYPTION" - Creates symmetric encryption keys   "RSA_SIGN_PSS_2048_SHA256" - RSASSA-PSS 2048 bit key with a SHA256 digest   "RSA_SIGN_PSS_3072_SHA256" - RSASSA-PSS 3072 bit key with a SHA256 digest   "RSA_SIGN_PSS_4096_SHA256" - RSASSA-PSS 4096 bit key with a SHA256 digest   "RSA_SIGN_PSS_4096_SHA512" - RSASSA-PSS 4096 bit key with a SHA512 digest   "RSA_SIGN_PKCS1_2048_SHA256" - RSASSA-PKCS1-v1_5 with a 2048 bit key and a SHA256 digest   "RSA_SIGN_PKCS1_3072_SHA256" - RSASSA-PKCS1-v1_5 with a 3072 bit key and a SHA256 digest   "RSA_SIGN_PKCS1_4096_SHA256" - RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA256 digest   "RSA_SIGN_PKCS1_4096_SHA512" - RSASSA-PKCS1-v1_5 with a 4096 bit key and a SHA512 digest   "RSA_DECRYPT_OAEP_2048_SHA256" - RSAES-OAEP 2048 bit key with a SHA256 digest   "RSA_DECRYPT_OAEP_3072_SHA256" - RSAES-OAEP 3072 bit key with a SHA256 digest   "RSA_DECRYPT_OAEP_4096_SHA256" - RSAES-OAEP 4096 bit key with a SHA256 digest   "RSA_DECRYPT_OAEP_4096_SHA512" - RSAES-OAEP 4096 bit key with a SHA512 digest   "EC_SIGN_P256_SHA256" - ECDSA on the NIST P-256 curve with a SHA256 digest   "EC_SIGN_P384_SHA384" - ECDSA on the NIST P-384 curve with a SHA384 digest   "EXTERNAL_SYMMETRIC_ENCRYPTION" - Algorithm representing symmetric encryption by an external key manager|
|primary_attestation_cert_chains_cavium_certs|text[]|Cavium certificate chain corresponding to the attestation|
|primary_attestation_cert_chains_google_card_certs|text[]|Google card certificate chain corresponding to the attestation|
|primary_attestation_cert_chains_google_partition_certs|text[]|Google partition certificate chain corresponding to the attestation|
|primary_attestation_content|text|The attestation data provided by the HSM when the key operation was performed|
|primary_attestation_format|text|The format of the attestation data  Possible values:   "ATTESTATION_FORMAT_UNSPECIFIED" - Not specified   "CAVIUM_V1_COMPRESSED" - Cavium HSM attestation compressed with gzip Note that this format is defined by Cavium and subject to change at any time   "CAVIUM_V2_COMPRESSED" - Cavium HSM attestation V2 compressed with gzip This is a new format introduced in Cavium's version 32-08|
|primary_create_time|timestamp without time zone|The time at which this CryptoKeyVersion was created|
|primary_destroy_event_time|timestamp without time zone|The time this CryptoKeyVersion's key material was destroyed Only present if state is DESTROYED|
|primary_destroy_time|timestamp without time zone|The time this CryptoKeyVersion's key material is scheduled for destruction Only present if state is DESTROY_SCHEDULED|
|primary_external_protection_level_options_external_key_uri|text|The URI for an external resource that this CryptoKeyVersion represents|
|primary_generate_time|timestamp without time zone|The time this CryptoKeyVersion's key material was generated|
|primary_import_failure_reason|text|The root cause of an import failure Only present if state is IMPORT_FAILED|
|primary_import_job|text|The name of the ImportJob used to import this CryptoKeyVersion Only present if the underlying key material was imported|
|primary_import_time|timestamp without time zone|The time at which this CryptoKeyVersion's key material was imported|
|primary_name|text|The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`|
|primary_protection_level|text|The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion|
|primary_state|text|The current state of the CryptoKeyVersion|
|purpose|text|Immutable The immutable purpose of this CryptoKey|
|rotation_period|text|next_rotation_time will be advanced by this period when the service automatically rotates a key Must be at least 24 hours and at most 876,000 hours If rotation_period is set, next_rotation_time must also be set Keys with purpose ENCRYPT_DECRYPT support automatic rotation For other keys, this field must be omitted|
|version_template_algorithm|text|Algorithm to use when creating a CryptoKeyVersion based on this template|
|version_template_protection_level|text|ProtectionLevel to use when creating a CryptoKeyVersion based on this template Immutable Defaults to SOFTWARE|
