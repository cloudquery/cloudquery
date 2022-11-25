# Table: aws_backup_plan_selections

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupSelection.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_backup_plans](aws_backup_plans.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|plan_arn|String|
|backup_plan_id|String|
|backup_selection|JSON|
|creation_date|Timestamp|
|creator_request_id|String|
|selection_id|String|
|result_metadata|JSON|