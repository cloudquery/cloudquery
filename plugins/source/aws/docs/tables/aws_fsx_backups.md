# Table: aws_fsx_backups

This table shows data for Amazon FSx Backups.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Backup.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|id|`utf8`|
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