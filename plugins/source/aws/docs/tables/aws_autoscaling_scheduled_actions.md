# Table: aws_autoscaling_scheduled_actions

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ScheduledUpdateGroupAction.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|auto_scaling_group_name|String|
|desired_capacity|Int|
|end_time|Timestamp|
|max_size|Int|
|min_size|Int|
|recurrence|String|
|scheduled_action_arn|String|
|scheduled_action_name|String|
|start_time|Timestamp|
|time|Timestamp|
|time_zone|String|