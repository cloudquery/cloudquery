# Table: aws_fsx_data_repository_associations

This table shows data for Amazon FSx Data Repository Associations.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryAssociation.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|association_id|`utf8`|
|batch_import_meta_data_on_create|`bool`|
|creation_time|`timestamp[us, tz=UTC]`|
|data_repository_path|`utf8`|
|data_repository_subdirectories|`list<item: utf8, nullable>`|
|failure_details|`json`|
|file_cache_id|`utf8`|
|file_cache_path|`utf8`|
|file_system_id|`utf8`|
|file_system_path|`utf8`|
|imported_file_chunk_size|`int64`|
|lifecycle|`utf8`|
|nfs|`json`|
|resource_arn|`utf8`|
|s3|`json`|