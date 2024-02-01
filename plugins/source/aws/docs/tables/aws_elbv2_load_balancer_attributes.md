# Table: aws_elbv2_load_balancer_attributes

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Load Balancer Attributes.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancerAttribute.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**load_balancer_arn**, **key**).
## Relations

This table depends on [aws_elbv2_load_balancers](aws_elbv2_load_balancers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancer_arn|`utf8`|
|key|`utf8`|
|value|`utf8`|