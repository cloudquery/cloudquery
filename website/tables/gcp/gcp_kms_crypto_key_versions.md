# Table: gcp_kms_crypto_key_versions

This table shows data for GCP Cloud Key Management Service (KMS) Crypto Key Versions.

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_crypto_keys](gcp_kms_crypto_keys).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|state|`utf8`|
|protection_level|`utf8`|
|algorithm|`utf8`|
|attestation|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|generate_time|`timestamp[us, tz=UTC]`|
|destroy_time|`timestamp[us, tz=UTC]`|
|destroy_event_time|`timestamp[us, tz=UTC]`|
|import_job|`utf8`|
|import_time|`timestamp[us, tz=UTC]`|
|import_failure_reason|`utf8`|
|generation_failure_reason|`utf8`|
|external_destruction_failure_reason|`utf8`|
|external_protection_level_options|`json`|
|reimport_eligible|`bool`|