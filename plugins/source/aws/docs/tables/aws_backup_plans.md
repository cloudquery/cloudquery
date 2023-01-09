# Table: aws_backup_plans

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_backup_plans:
  - [aws_backup_plan_selections](aws_backup_plan_selections.md)

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
|advanced_backup_settings|JSON|
|backup_plan|JSON|
|backup_plan_arn|String|
|backup_plan_id|String|
|creation_date|Timestamp|
|creator_request_id|String|
|deletion_date|Timestamp|
|last_execution_date|Timestamp|
|version_id|String|
|result_metadata|JSON|