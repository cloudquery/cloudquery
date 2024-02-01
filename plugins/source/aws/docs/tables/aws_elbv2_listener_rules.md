# Table: aws_elbv2_listener_rules

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Listener Rules.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Rule.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_elbv2_listeners](aws_elbv2_listeners.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|listener_arn|`utf8`|
|arn|`utf8`|
|actions|`json`|
|conditions|`json`|
|is_default|`bool`|
|priority|`utf8`|
|rule_arn|`utf8`|