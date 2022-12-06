# Table: aws_backup_global_settings

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeGlobalSettings.html

The primary key for this table is **account_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|global_settings|JSON|
|last_update_time|Timestamp|
|result_metadata|JSON|