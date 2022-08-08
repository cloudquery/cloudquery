
# Table: aws_autoscaling_group_lifecycle_hooks
Describes a lifecycle hook, which tells Amazon EC2 Auto Scaling that you want to perform an action whenever it launches instances or terminates instances.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_cq_id|uuid|Unique CloudQuery ID of aws_autoscaling_groups table (FK)|
|auto_scaling_group_name|text|The name of the Auto Scaling group for the lifecycle hook.|
|default_result|text|Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs|
|global_timeout|integer|The maximum time, in seconds, that an instance can remain in a Pending:Wait or Terminating:Wait state|
|heartbeat_timeout|integer|The maximum time, in seconds, that can elapse before the lifecycle hook times out|
|lifecycle_hook_name|text|The name of the lifecycle hook.|
|lifecycle_transition|text|The state of the EC2 instance to which to attach the lifecycle hook|
|notification_metadata|text|Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.|
|notification_target_arn|text|The ARN of the target that Amazon EC2 Auto Scaling sends notifications to when an instance is in the transition state for the lifecycle hook|
|role_arn|text|The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.|
