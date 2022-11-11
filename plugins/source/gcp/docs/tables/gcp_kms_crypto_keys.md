# Table: gcp_kms_crypto_keys



The primary key for this table is **_cq_id**.

## Relations
This table depends on [gcp_kms_keyrings](gcp_kms_keyrings.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|create_time|String|
|crypto_key_backend|String|
|destroy_scheduled_duration|String|
|import_only|Bool|
|labels|JSON|
|name|String|
|next_rotation_time|String|
|primary|JSON|
|purpose|String|
|rotation_period|String|
|version_template|JSON|