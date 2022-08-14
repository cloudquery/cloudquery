
# Table: aws_elbv2_listener_default_action_forward_config_target_groups
Information about how traffic will be distributed between multiple target groups in a forward rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|listener_default_action_cq_id|uuid|Unique CloudQuery ID of aws_elbv2_listener_default_actions table (FK)|
|target_group_arn|text|The Amazon Resource Name (ARN) of the target group.|
|weight|integer|The weight|
