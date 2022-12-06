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
|rotation_period|Int|
|name|String|
|primary|JSON|
|purpose|String|
|create_time|Timestamp|
|next_rotation_time|Timestamp|
|version_template|JSON|
|labels|JSON|
|import_only|Bool|
|destroy_scheduled_duration|Int|
|crypto_key_backend|String|