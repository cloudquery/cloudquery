# Table: aws_elbv2_target_groups

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_target_groups:
  - [aws_elbv2_target_group_target_health_descriptions](aws_elbv2_target_group_target_health_descriptions.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|health_check_enabled|Bool|
|health_check_interval_seconds|Int|
|health_check_path|String|
|health_check_port|String|
|health_check_protocol|String|
|health_check_timeout_seconds|Int|
|healthy_threshold_count|Int|
|ip_address_type|String|
|load_balancer_arns|StringArray|
|matcher|JSON|
|port|Int|
|protocol|String|
|protocol_version|String|
|target_group_name|String|
|target_type|String|
|unhealthy_threshold_count|Int|
|vpc_id|String|