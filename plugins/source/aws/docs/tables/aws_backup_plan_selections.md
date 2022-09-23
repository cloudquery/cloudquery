# Table: aws_backup_plan_selections


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_backup_plans`](aws_backup_plans.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|plan_arn|String|
|backup_plan_id|String|
|backup_selection|JSON|
|creation_date|Timestamp|
|creator_request_id|String|
|selection_id|String|
|result_metadata|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|