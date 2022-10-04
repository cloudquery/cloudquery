# Table: aws_elbv2_load_balancer_attributes



The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_elbv2_load_balancers`](aws_elbv2_load_balancers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|load_balancer_arn|String|
|key|String|
|value|String|