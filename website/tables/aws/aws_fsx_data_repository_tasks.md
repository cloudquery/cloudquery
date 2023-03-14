# Table: aws_fsx_data_repository_tasks

This table shows data for Amazon FSx Data Repository Tasks.

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
|tags|JSON|
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
|resource_arn|String|
|start_time|Timestamp|
|status|JSON|