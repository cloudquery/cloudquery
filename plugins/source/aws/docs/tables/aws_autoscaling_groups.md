# Table: aws_autoscaling_groups

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_autoscaling_groups:
  - [aws_autoscaling_group_scaling_policies](aws_autoscaling_group_scaling_policies.md)
  - [aws_autoscaling_group_lifecycle_hooks](aws_autoscaling_group_lifecycle_hooks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|load_balancers|JSON|
|load_balancer_target_groups|JSON|
|arn (PK)|String|
|auto_scaling_group_name|String|
|availability_zones|StringArray|
|created_time|Timestamp|
|default_cooldown|Int|
|desired_capacity|Int|
|health_check_type|String|
|max_size|Int|
|min_size|Int|
|auto_scaling_group_arn|String|
|capacity_rebalance|Bool|
|context|String|
|default_instance_warmup|Int|
|desired_capacity_type|String|
|enabled_metrics|JSON|
|health_check_grace_period|Int|
|instances|JSON|
|launch_configuration_name|String|
|launch_template|JSON|
|load_balancer_names|StringArray|
|max_instance_lifetime|Int|
|mixed_instances_policy|JSON|
|new_instances_protected_from_scale_in|Bool|
|placement_group|String|
|predicted_capacity|Int|
|service_linked_role_arn|String|
|status|String|
|suspended_processes|JSON|
|tags|JSON|
|target_group_ar_ns|StringArray|
|termination_policies|StringArray|
|traffic_sources|JSON|
|vpc_zone_identifier|String|
|warm_pool_configuration|JSON|
|warm_pool_size|Int|
|notification_configurations|JSON|