# Table: aws_fsx_data_repository_associations



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
|association_id|String|
|batch_import_meta_data_on_create|Bool|
|creation_time|Timestamp|
|data_repository_path|String|
|failure_details|JSON|
|file_system_id|String|
|file_system_path|String|
|imported_file_chunk_size|Int|
|lifecycle|String|
|s3|JSON|