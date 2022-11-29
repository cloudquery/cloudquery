# Table: aws_autoscaling_group_lifecycle_hooks

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_autoscaling_groups](aws_autoscaling_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|group_arn|String|
|auto_scaling_group_name|String|
|default_result|String|
|global_timeout|Int|
|heartbeat_timeout|Int|
|lifecycle_hook_name|String|
|lifecycle_transition|String|
|notification_metadata|String|
|notification_target_arn|String|
|role_arn|String|