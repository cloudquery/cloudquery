
# Table: aws_elbv2_target_group_target_health_descriptions
Information about the health of a target.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|target_group_cq_id|uuid|Unique CloudQuery ID of aws_elbv2_target_groups table (FK)|
|health_check_port|text|The port to use to connect with the target.|
|target_id|text|The ID of the target.|
|target_availability_zone|text|An Availability Zone or all.|
|target_port|integer|The port on which the target is listening.|
|target_health_description|text|A description of the target health that provides additional details.|
|target_health_reason|text|The reason code. If the target state is healthy, a reason code is not provided.|
|target_health_state|text|The state of the target.|
