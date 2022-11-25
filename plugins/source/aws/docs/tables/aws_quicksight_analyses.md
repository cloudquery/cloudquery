# Table: aws_quicksight_analyses



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
|tags|JSON|
|arn (PK)|String|
|analysis_id|String|
|created_time|Timestamp|
|data_set_arns|StringArray|
|errors|JSON|
|last_updated_time|Timestamp|
|name|String|
|sheets|JSON|
|status|String|
|theme_arn|String|