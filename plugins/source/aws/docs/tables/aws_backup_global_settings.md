# Table: aws_backup_global_settings

This table shows data for Backup Global Settings.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeGlobalSettings.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|global_settings|`json`|
|last_update_time|`timestamp[us, tz=UTC]`|
|result_metadata|`json`|