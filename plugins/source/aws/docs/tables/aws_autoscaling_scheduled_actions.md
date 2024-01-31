# Table: aws_autoscaling_scheduled_actions

This table shows data for Auto Scaling Scheduled Actions.

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ScheduledUpdateGroupAction.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|auto_scaling_group_name|`utf8`|
|desired_capacity|`int64`|
|end_time|`timestamp[us, tz=UTC]`|
|max_size|`int64`|
|min_size|`int64`|
|recurrence|`utf8`|
|scheduled_action_arn|`utf8`|
|scheduled_action_name|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|time|`timestamp[us, tz=UTC]`|
|time_zone|`utf8`|