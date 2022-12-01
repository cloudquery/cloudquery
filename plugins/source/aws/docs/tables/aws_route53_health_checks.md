# Table: aws_route53_health_checks

https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheck.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|cloud_watch_alarm_configuration_dimensions|JSON|
|caller_reference|String|
|health_check_config|JSON|
|health_check_version|Int|
|id|String|
|cloud_watch_alarm_configuration|JSON|
|linked_service|JSON|