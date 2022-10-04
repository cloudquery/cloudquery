# Table: aws_fsx_data_repository_tasks



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|creation_time|Timestamp|
|file_system_id|String|
|lifecycle|String|
|task_id|String|
|type|String|
|end_time|Timestamp|
|failure_details|JSON|
|paths|StringArray|
|report|JSON|
|start_time|Timestamp|
|status|JSON|