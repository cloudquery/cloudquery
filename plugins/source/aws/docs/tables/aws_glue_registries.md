# Table: aws_glue_registries



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_glue_registries:
  - [aws_glue_registry_schemas](aws_glue_registry_schemas.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|created_time|String|
|description|String|
|registry_name|String|
|status|String|
|updated_time|String|