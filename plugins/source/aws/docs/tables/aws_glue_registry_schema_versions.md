# Table: aws_glue_registry_schema_versions


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|