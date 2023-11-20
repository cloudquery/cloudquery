# Table: aws_autoscaling_group_lifecycle_hooks

This table shows data for Auto Scaling Group Lifecycle Hooks.

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html

The composite primary key for this table is (**group_arn**, **lifecycle_hook_name**).

## Relations

This table depends on [aws_autoscaling_groups](aws_autoscaling_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|group_arn (PK)|`utf8`|
|auto_scaling_group_name|`utf8`|
|default_result|`utf8`|
|global_timeout|`int64`|
|heartbeat_timeout|`int64`|
|lifecycle_hook_name (PK)|`utf8`|
|lifecycle_transition|`utf8`|
|notification_metadata|`utf8`|
|notification_target_arn|`utf8`|
|role_arn|`utf8`|