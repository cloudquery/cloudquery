# Table: aws_autoscaling_group_lifecycle_hooks

This table shows data for Auto Scaling Group Lifecycle Hooks.

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_autoscaling_groups](aws_autoscaling_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|group_arn|`utf8`|
|auto_scaling_group_name|`utf8`|
|default_result|`utf8`|
|global_timeout|`int64`|
|heartbeat_timeout|`int64`|
|lifecycle_hook_name|`utf8`|
|lifecycle_transition|`utf8`|
|notification_metadata|`utf8`|
|notification_target_arn|`utf8`|
|role_arn|`utf8`|