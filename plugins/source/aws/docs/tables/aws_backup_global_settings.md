# Table: aws_backup_global_settings

This table shows data for Backup Global Settings.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeGlobalSettings.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|global_settings|`json`|
|last_update_time|`timestamp[us, tz=UTC]`|