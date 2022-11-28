# Table: aws_fsx_data_repository_tasks

https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryTask.html

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
|creation_time|Timestamp|
|lifecycle|String|
|task_id|String|
|type|String|
|capacity_to_release|Int|
|end_time|Timestamp|
|failure_details|JSON|
|file_cache_id|String|
|file_system_id|String|
|paths|StringArray|
|report|JSON|
|start_time|Timestamp|
|status|JSON|
|tags|JSON|