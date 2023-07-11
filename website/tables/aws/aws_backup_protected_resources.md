# Table: aws_backup_protected_resources

This table shows data for Backup Protected Resources.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListProtectedResources.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|last_backup_time|`timestamp[us, tz=UTC]`|
|resource_arn|`utf8`|
|resource_name|`utf8`|
|resource_type|`utf8`|