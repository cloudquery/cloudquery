# Table: aws_glue_registry_schemas



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_glue_registries](aws_glue_registries.md).

The following tables depend on aws_glue_registry_schemas:
  - [aws_glue_registry_schema_versions](aws_glue_registry_schema_versions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|tags|JSON|
|compatibility|String|
|created_time|String|
|data_format|String|
|description|String|
|latest_schema_version|Int|
|next_schema_version|Int|
|registry_arn|String|
|registry_name|String|
|schema_checkpoint|Int|
|schema_name|String|
|schema_status|String|
|updated_time|String|
|result_metadata|JSON|