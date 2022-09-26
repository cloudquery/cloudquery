# Table: aws_elbv2_target_group_target_health_descriptions


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_elbv2_target_groups`](aws_elbv2_target_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|target_group_arn|String|
|health_check_port|String|
|target|JSON|
|target_health|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|