
# Table: aws_autoscaling_groups
Describes an Auto Scaling group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|load_balancers|jsonb||
|load_balancer_target_groups|jsonb||
|notifications_configurations|jsonb||
|name|text|The name of the Auto Scaling group.|
|availability_zones|text[]|One or more Availability Zones for the group.|
|created_time|timestamp without time zone|The date and time the group was created.|
|default_cooldown|integer|The duration of the default cooldown period, in seconds.|
|desired_capacity|integer|The desired size of the group.|
|health_check_type|text|The service to use for the health checks|
|max_size|integer|The maximum size of the group.|
|min_size|integer|The minimum size of the group.|
|arn|text|The Amazon Resource Name (ARN) of the Auto Scaling group.|
|capacity_rebalance|boolean|Indicates whether Capacity Rebalancing is enabled.|
|enabled_metrics|jsonb|The metrics enabled for the group.|
|health_check_grace_period|integer|The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.|
|launch_configuration_name|text|The name of the associated launch configuration.|
|launch_template_id|text|The ID of the launch template|
|launch_template_name|text|The name of the launch template|
|launch_template_version|text|The version number, $Latest, or $Default|
|load_balancer_names|text[]|One or more load balancers associated with the group.|
|max_instance_lifetime|integer|The maximum amount of time, in seconds, that an instance can be in service. Valid Range: Minimum value of 0.|
|mixed_instances_policy|jsonb|The mixed instances policy for the group.|
|new_instances_protected_from_scale_in|boolean|Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.|
|placement_group|text|The name of the placement group into which to launch your instances, if any.|
|service_linked_role_arn|text|The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.|
|status|text|The current state of the group when the DeleteAutoScalingGroup operation is in progress.|
|suspended_processes|jsonb|The suspended processes associated with the group.|
|target_group_arns|text[]|The Amazon Resource Names (ARN) of the target groups for your load balancer.|
|termination_policies|text[]|The termination policies for the group.|
|vpc_zone_identifier|text|One or more subnet IDs, if applicable, separated by commas.|
