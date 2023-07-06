# Table: aws_fsx_data_repository_tasks

This table shows data for Amazon FSx Data Repository Tasks.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryTask.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|lifecycle|`utf8`|
|task_id|`utf8`|
|type|`utf8`|
|capacity_to_release|`int64`|
|end_time|`timestamp[us, tz=UTC]`|
|failure_details|`json`|
|file_cache_id|`utf8`|
|file_system_id|`utf8`|
|paths|`list<item: utf8, nullable>`|
|report|`json`|
|resource_arn|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|status|`json`|