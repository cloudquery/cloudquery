# Table: aws_fsx_backups

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Backup.html

The composite primary key for this table is (**account_id**, **region**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|id (PK)|String|
|creation_time|Timestamp|
|file_system|JSON|
|lifecycle|String|
|type|String|
|directory_information|JSON|
|failure_details|JSON|
|kms_key_id|String|
|owner_id|String|
|progress_percent|Int|
|resource_arn|String|
|resource_type|String|
|source_backup_id|String|
|source_backup_region|String|
|tags|JSON|
|volume|JSON|