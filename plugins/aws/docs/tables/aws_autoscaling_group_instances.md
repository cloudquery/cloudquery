
# Table: aws_autoscaling_group_instances
Describes an EC2 instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_cq_id|uuid|Unique CloudQuery ID of aws_autoscaling_groups table (FK)|
|availability_zone|text|The Availability Zone in which the instance is running.|
|health_status|text|The last reported health status of the instance|
|id|text|The ID of the instance.|
|lifecycle_state|text|A description of the current lifecycle state|
|protected_from_scale_in|boolean|Indicates whether the instance is protected from termination by Amazon EC2 Auto Scaling when scaling in.|
|type|text|The instance type of the EC2 instance.|
|launch_configuration_name|text|The launch configuration associated with the instance.|
|launch_template_id|text|The ID of the launch template|
|launch_template_name|text|The name of the launch template|
|launch_template_version|text|The version number, $Latest, or $Default|
|weighted_capacity|text|The number of capacity units contributed by the instance based on its instance type|
