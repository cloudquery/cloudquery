# Table: aws_cloudhsmv2_backups

https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Backup.html

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
|backup_id|String|
|backup_state|String|
|cluster_id|String|
|copy_timestamp|Timestamp|
|create_timestamp|Timestamp|
|delete_timestamp|Timestamp|
|never_expires|Bool|
|source_backup|String|
|source_cluster|String|
|source_region|String|
|tag_list|JSON|