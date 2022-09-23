# Table: aws_elbv1_load_balancer_policies


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_elbv1_load_balancers`](aws_elbv1_load_balancers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|load_balancer_arn|String|
|load_balancer_name|String|
|policy_attribute_descriptions|JSON|
|policy_name|String|
|policy_type_name|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|