# Table: aws_docdb_cluster_parameters

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Parameter.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_docdb_engine_versions](aws_docdb_engine_versions.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|allowed_values|String|
|apply_method|String|
|apply_type|String|
|data_type|String|
|description|String|
|is_modifiable|Bool|
|minimum_engine_version|String|
|parameter_name|String|
|parameter_value|String|
|source|String|