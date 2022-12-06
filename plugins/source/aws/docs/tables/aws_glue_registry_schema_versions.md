# Table: aws_glue_registry_schema_versions



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_glue_registry_schemas](aws_glue_registry_schemas.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|registry_schema_arn|String|
|metadata|JSON|
|created_time|String|
|data_format|String|
|schema_arn|String|
|schema_definition|String|
|schema_version_id|String|
|status|String|
|version_number|Int|
|result_metadata|JSON|