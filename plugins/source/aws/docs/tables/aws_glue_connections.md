# Table: aws_glue_connections



The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|connection_properties|JSON|
|connection_type|String|
|creation_time|Timestamp|
|description|String|
|last_updated_by|String|
|last_updated_time|Timestamp|
|match_criteria|StringArray|
|name|String|
|physical_connection_requirements|JSON|