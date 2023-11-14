# Table: aws_cloudhsmv2_backups

This table shows data for AWS CloudHSM v2 Backups.

https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Backup.html

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
|backup_id|`utf8`|
|backup_state|`utf8`|
|cluster_id|`utf8`|
|copy_timestamp|`timestamp[us, tz=UTC]`|
|create_timestamp|`timestamp[us, tz=UTC]`|
|delete_timestamp|`timestamp[us, tz=UTC]`|
|never_expires|`bool`|
|source_backup|`utf8`|
|source_cluster|`utf8`|
|source_region|`utf8`|