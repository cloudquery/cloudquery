# Table: aws_fsx_data_repository_associations

https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryAssociation.html

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
|association_id|String|
|batch_import_meta_data_on_create|Bool|
|creation_time|Timestamp|
|data_repository_path|String|
|data_repository_subdirectories|StringArray|
|failure_details|JSON|
|file_cache_id|String|
|file_cache_path|String|
|file_system_id|String|
|file_system_path|String|
|imported_file_chunk_size|Int|
|lifecycle|String|
|nfs|JSON|
|s3|JSON|
|tags|JSON|