# Table: aws_backup_plan_selections

This table shows data for Backup Plan Selections.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupSelection.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_backup_plans](aws_backup_plans).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|plan_arn|`utf8`|
|backup_plan_id|`utf8`|
|backup_selection|`json`|
|creation_date|`timestamp[us, tz=UTC]`|
|creator_request_id|`utf8`|
|selection_id|`utf8`|
|result_metadata|`json`|