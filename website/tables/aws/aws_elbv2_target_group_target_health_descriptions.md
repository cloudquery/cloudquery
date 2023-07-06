# Table: aws_elbv2_target_group_target_health_descriptions

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Target Group Target Health Descriptions.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetHealthDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_elbv2_target_groups](aws_elbv2_target_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|target_group_arn|`utf8`|
|health_check_port|`utf8`|
|target|`json`|
|target_health|`json`|