# Table: aws_fsx_data_repository_tasks


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|