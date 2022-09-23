# Table: aws_fsx_data_repository_associations


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|