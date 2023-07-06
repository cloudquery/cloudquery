# Table: gcp_kms_crypto_keys

This table shows data for GCP Cloud Key Management Service (KMS) Crypto Keys.

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKey

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_keyrings](gcp_kms_keyrings).

The following tables depend on gcp_kms_crypto_keys:
  - [gcp_kms_crypto_key_versions](gcp_kms_crypto_key_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|rotation_period|`int64`|
|name (PK)|`utf8`|
|primary|`json`|
|purpose|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|next_rotation_time|`timestamp[us, tz=UTC]`|
|version_template|`json`|
|labels|`json`|
|import_only|`bool`|
|destroy_scheduled_duration|`int64`|
|crypto_key_backend|`utf8`|