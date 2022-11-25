# Table: aws_elbv1_load_balancer_policies

https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_elbv1_load_balancers](aws_elbv1_load_balancers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|load_balancer_arn|String|
|load_balancer_name|String|
|policy_attribute_descriptions|JSON|
|policy_name|String|
|policy_type_name|String|