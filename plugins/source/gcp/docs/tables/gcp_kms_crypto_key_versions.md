# Table: gcp_kms_crypto_key_versions

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_crypto_keys](gcp_kms_crypto_keys.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|state|String|
|protection_level|String|
|algorithm|String|
|attestation|JSON|
|create_time|Timestamp|
|generate_time|Timestamp|
|destroy_time|Timestamp|
|destroy_event_time|Timestamp|
|import_job|String|
|import_time|Timestamp|
|import_failure_reason|String|
|external_protection_level_options|JSON|
|reimport_eligible|Bool|