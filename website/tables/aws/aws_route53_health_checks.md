# Table: aws_route53_health_checks

This table shows data for Amazon Route 53 Health Checks.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheck.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|cloud_watch_alarm_configuration_dimensions|`json`|
|caller_reference|`utf8`|
|health_check_config|`json`|
|health_check_version|`int64`|
|id|`utf8`|
|cloud_watch_alarm_configuration|`json`|
|linked_service|`json`|