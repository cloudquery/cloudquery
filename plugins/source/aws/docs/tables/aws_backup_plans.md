# Table: aws_backup_plans

This table shows data for Backup Plans.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**arn**, **version_id**).
## Relations

The following tables depend on aws_backup_plans:
  - [aws_backup_plan_selections](aws_backup_plan_selections.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|advanced_backup_settings|`json`|
|backup_plan|`json`|
|backup_plan_arn|`utf8`|
|backup_plan_id|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|creator_request_id|`utf8`|
|deletion_date|`timestamp[us, tz=UTC]`|
|last_execution_date|`timestamp[us, tz=UTC]`|
|version_id|`utf8`|