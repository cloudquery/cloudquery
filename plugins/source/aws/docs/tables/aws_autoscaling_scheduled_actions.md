
# Table: aws_autoscaling_scheduled_actions
Describes a scheduled scaling action.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|auto_scaling_group_name|text|The name of the Auto Scaling group.|
|desired_capacity|integer|The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.|
|end_time|timestamp without time zone|The date and time in UTC for the recurring schedule to end|
|max_size|integer|The maximum size of the Auto Scaling group.|
|min_size|integer|The minimum size of the Auto Scaling group.|
|recurrence|text|The recurring schedule for the action, in Unix cron syntax format|
|arn|text|The Amazon Resource Name (ARN) of the scheduled action.|
|name|text|The name of the scheduled action.|
|start_time|timestamp without time zone|The date and time in UTC for this action to start|
|time|timestamp without time zone|This parameter is no longer used.|
|time_zone|text|The time zone for the cron expression.|
