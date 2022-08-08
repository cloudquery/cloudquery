
# Table: aws_autoscaling_group_tags
Describes a tag for an Auto Scaling group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_cq_id|uuid|Unique CloudQuery ID of aws_autoscaling_groups table (FK)|
|key|text|The tag key.|
|propagate_at_launch|boolean|Determines whether the tag is added to new instances as they are launched in the group.|
|resource_id|text|The name of the group.|
|resource_type|text|The type of resource|
|value|text|The tag value.|
