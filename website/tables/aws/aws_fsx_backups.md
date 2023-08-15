# Table: aws_fsx_backups

This table shows data for Amazon FSx Backups.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Backup.html

The composite primary key for this table is (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|id (PK)|`utf8`|
|tags|`json`|
|backup_id|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|file_system|`json`|
|lifecycle|`utf8`|
|type|`utf8`|
|directory_information|`json`|
|failure_details|`json`|
|kms_key_id|`utf8`|
|owner_id|`utf8`|
|progress_percent|`int64`|
|resource_arn|`utf8`|
|resource_type|`utf8`|
|source_backup_id|`utf8`|
|source_backup_region|`utf8`|
|volume|`json`|