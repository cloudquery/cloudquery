# Table: aws_computeoptimizer_autoscaling_group_recommendations

This table shows data for Compute Optimizer Auto Scaling Group Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_AutoScalingGroupRecommendation.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **auto_scaling_group_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|auto_scaling_group_arn|`utf8`|
|auto_scaling_group_name|`utf8`|
|current_configuration|`json`|
|current_instance_gpu_info|`json`|
|current_performance_risk|`utf8`|
|effective_recommendation_preferences|`json`|
|finding|`utf8`|
|inferred_workload_types|`list<item: utf8, nullable>`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|look_back_period_in_days|`float64`|
|recommendation_options|`json`|
|utilization_metrics|`json`|