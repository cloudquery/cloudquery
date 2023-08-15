# Table: aws_autoscaling_groups

This table shows data for Auto Scaling Groups.

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_autoscaling_groups:
  - [aws_autoscaling_group_lifecycle_hooks](aws_autoscaling_group_lifecycle_hooks)
  - [aws_autoscaling_group_scaling_policies](aws_autoscaling_group_scaling_policies)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancers|`json`|
|load_balancer_target_groups|`json`|
|arn (PK)|`utf8`|
|tags|`json`|
|tags_raw|`json`|
|auto_scaling_group_name|`utf8`|
|availability_zones|`list<item: utf8, nullable>`|
|created_time|`timestamp[us, tz=UTC]`|
|default_cooldown|`int64`|
|desired_capacity|`int64`|
|health_check_type|`utf8`|
|max_size|`int64`|
|min_size|`int64`|
|auto_scaling_group_arn|`utf8`|
|capacity_rebalance|`bool`|
|context|`utf8`|
|default_instance_warmup|`int64`|
|desired_capacity_type|`utf8`|
|enabled_metrics|`json`|
|health_check_grace_period|`int64`|
|instances|`json`|
|launch_configuration_name|`utf8`|
|launch_template|`json`|
|load_balancer_names|`list<item: utf8, nullable>`|
|max_instance_lifetime|`int64`|
|mixed_instances_policy|`json`|
|new_instances_protected_from_scale_in|`bool`|
|placement_group|`utf8`|
|predicted_capacity|`int64`|
|service_linked_role_arn|`utf8`|
|status|`utf8`|
|suspended_processes|`json`|
|target_group_arns|`list<item: utf8, nullable>`|
|termination_policies|`list<item: utf8, nullable>`|
|traffic_sources|`json`|
|vpc_zone_identifier|`utf8`|
|warm_pool_configuration|`json`|
|warm_pool_size|`int64`|
|notification_configurations|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Auto Scaling groups associated with a load balancer should use health checks

```sql
SELECT
  'Auto Scaling groups associated with a load balancer should use health checks'
    AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN array_length(load_balancer_names, 1) > 0
  AND health_check_type IS DISTINCT FROM 'ELB'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_autoscaling_groups;
```


