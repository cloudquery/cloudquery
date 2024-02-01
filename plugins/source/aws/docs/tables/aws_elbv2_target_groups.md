# Table: aws_elbv2_target_groups

This table shows data for Amazon Elastic Load Balancer (ELB) v2 Target Groups.

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_elbv2_target_groups:
  - [aws_elbv2_target_group_target_health_descriptions](aws_elbv2_target_group_target_health_descriptions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|health_check_enabled|`bool`|
|health_check_interval_seconds|`int64`|
|health_check_path|`utf8`|
|health_check_port|`utf8`|
|health_check_protocol|`utf8`|
|health_check_timeout_seconds|`int64`|
|healthy_threshold_count|`int64`|
|ip_address_type|`utf8`|
|load_balancer_arns|`list<item: utf8, nullable>`|
|matcher|`json`|
|port|`int64`|
|protocol|`utf8`|
|protocol_version|`utf8`|
|target_group_arn|`utf8`|
|target_group_name|`utf8`|
|target_type|`utf8`|
|unhealthy_threshold_count|`int64`|
|vpc_id|`utf8`|