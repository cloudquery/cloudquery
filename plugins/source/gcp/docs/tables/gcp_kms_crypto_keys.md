# Table: gcp_kms_crypto_keys


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`gcp_kms_keyrings`](gcp_kms_keyrings.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name|String|
|primary|JSON|
|purpose|Int|
|create_time|Timestamp|
|next_rotation_time|Timestamp|
|version_template|JSON|
|labels|JSON|
|import_only|Bool|
|destroy_scheduled_duration|JSON|
|crypto_key_backend|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|