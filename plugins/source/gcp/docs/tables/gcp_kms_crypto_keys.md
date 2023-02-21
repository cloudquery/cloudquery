# Table: gcp_kms_crypto_keys

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKey

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_keyrings](gcp_kms_keyrings.md).

The following tables depend on gcp_kms_crypto_keys:
  - [gcp_kms_crypto_key_versions](gcp_kms_crypto_key_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|rotation_period|Int|
|name (PK)|String|
|primary|JSON|
|purpose|String|
|create_time|Timestamp|
|next_rotation_time|Timestamp|
|version_template|JSON|
|labels|JSON|
|import_only|Bool|
|destroy_scheduled_duration|Int|
|crypto_key_backend|String|