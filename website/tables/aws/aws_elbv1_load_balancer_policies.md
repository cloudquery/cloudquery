# Table: aws_elbv1_load_balancer_policies

This table shows data for Amazon Elastic Load Balancer (ELB) v1 Load Balancer Policies.

https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_elbv1_load_balancers](aws_elbv1_load_balancers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancer_arn|`utf8`|
|load_balancer_name|`utf8`|
|policy_attribute_descriptions|`json`|
|policy_name|`utf8`|
|policy_type_name|`utf8`|