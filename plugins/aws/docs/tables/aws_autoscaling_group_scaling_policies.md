
# Table: aws_autoscaling_group_scaling_policies
Describes a scaling policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_cq_id|uuid|Unique CloudQuery ID of aws_autoscaling_groups table (FK)|
|adjustment_type|text|Specifies how the scaling adjustment is interpreted (for example, an absolute number or a percentage)|
|alarms|jsonb|The CloudWatch alarms related to the policy.|
|auto_scaling_group_name|text|The name of the Auto Scaling group.|
|cooldown|integer|The duration of the policy's cooldown period, in seconds.|
|enabled|boolean|Indicates whether the policy is enabled (true) or disabled (false).|
|estimated_instance_warmup|integer|The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.|
|metric_aggregation_type|text|The aggregation type for the CloudWatch metrics|
|min_adjustment_magnitude|integer|The minimum value to scale by when the adjustment type is PercentChangeInCapacity.|
|min_adjustment_step|integer|Available for backward compatibility|
|arn|text|The Amazon Resource Name (ARN) of the policy.|
|name|text|The name of the scaling policy.|
|type|text|One of the following policy types:  * TargetTrackingScaling  * StepScaling  * SimpleScaling (default)  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html) and Step and simple scaling policies (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html) in the Amazon EC2 Auto Scaling User Guide.|
|scaling_adjustment|integer|The amount by which to scale, based on the specified adjustment type|
|step_adjustments|jsonb|A set of adjustments that enable you to scale based on the size of the alarm breach.|
|target_tracking_configuration_target_value|float|The target value for the metric.|
|target_tracking_configuration_customized_metric_name|text|The name of the metric.|
|target_tracking_configuration_customized_metric_namespace|text|The namespace of the metric.|
|target_tracking_configuration_customized_metric_statistic|text|The statistic of the metric.|
|target_tracking_configuration_customized_metric_dimensions|jsonb|The dimensions of the metric|
|target_tracking_configuration_customized_metric_unit|text|The unit of the metric.|
|target_tracking_configuration_disable_scale_in|boolean|Indicates whether scaling in by the target tracking scaling policy is disabled. If scaling in is disabled, the target tracking scaling policy doesn't remove instances from the Auto Scaling group|
|target_tracking_configuration_predefined_metric_type|text|The metric type|
|target_tracking_configuration_predefined_metric_resource_label|text|Identifies the resource associated with the metric type|
