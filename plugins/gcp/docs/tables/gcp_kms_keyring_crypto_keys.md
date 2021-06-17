
# Table: gcp_kms_keyring_crypto_keys
A CryptoKey represents a logical key that can be used for cryptographic operations.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|keyring_id|uuid|Unique ID of gcp_kms_keyrings table (FK)|
|project_id|text|GCP Project Id of the resource|
|location|text|Location of the resource|
|create_time|text|The time at which this CryptoKey was created|
|labels|jsonb|Labels with user-defined metadata For more information, see Labeling Keys (https://cloudgooglecom/kms/docs/labeling-keys)|
|name|text|The resource name for this CryptoKey in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*`|
|next_rotation_time|text|At next_rotation_time, the Key Management Service will automatically: 1 Create a new version of this CryptoKey 2 Mark the new version as primary Key rotations performed manually via CreateCryptoKeyVersion and UpdateCryptoKeyPrimaryVersion do not affect next_rotation_time Keys with purpose ENCRYPT_DECRYPT support automatic rotation For other keys, this field must be omitted|
|primary_algorithm|text|The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports|
|primary_attestation_cert_chains_cavium_certs|text[]|Cavium certificate chain corresponding to the attestation|
|primary_attestation_cert_chains_google_card_certs|text[]|Google card certificate chain corresponding to the attestation|
|primary_attestation_cert_chains_google_partition_certs|text[]|Google partition certificate chain corresponding to the attestation|
|primary_attestation_content|text|The attestation data provided by the HSM when the key operation was performed|
|primary_attestation_format|text|The format of the attestation data|
|primary_create_time|text|The time at which this CryptoKeyVersion was created|
|primary_destroy_event_time|text|The time this CryptoKeyVersion's key material was destroyed Only present if state is DESTROYED|
|primary_destroy_time|text|The time this CryptoKeyVersion's key material is scheduled for destruction Only present if state is DESTROY_SCHEDULED|
|primary_external_protection_level_options_external_key_uri|text|The URI for an external resource that this CryptoKeyVersion represents|
|primary_generate_time|text|The time this CryptoKeyVersion's key material was generated|
|primary_import_failure_reason|text|The root cause of an import failure Only present if state is IMPORT_FAILED|
|primary_import_job|text|The name of the ImportJob used to import this CryptoKeyVersion Only present if the underlying key material was imported|
|primary_import_time|text|The time at which this CryptoKeyVersion's key material was imported|
|primary_name|text|The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`|
|primary_protection_level|text|The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion|
|primary_state|text|The current state of the CryptoKeyVersion|
|purpose|text|Immutable The immutable purpose of this CryptoKey|
|rotation_period|text|next_rotation_time will be advanced by this period when the service automatically rotates a key Must be at least 24 hours and at most 876,000 hours If rotation_period is set, next_rotation_time must also be set Keys with purpose ENCRYPT_DECRYPT support automatic rotation For other keys, this field must be omitted|
|version_template_algorithm|text|Required Algorithm to use when creating a CryptoKeyVersion based on this template For backwards compatibility, GOOGLE_SYMMETRIC_ENCRYPTION is implied if both this field is omitted and CryptoKeypurpose is ENCRYPT_DECRYPT|
|version_template_protection_level|text|ProtectionLevel to use when creating a CryptoKeyVersion based on this template Immutable Defaults to SOFTWARE|
