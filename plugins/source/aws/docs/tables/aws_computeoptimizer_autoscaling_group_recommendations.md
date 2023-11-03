# Table: aws_computeoptimizer_autoscaling_group_recommendations

This table shows data for Compute Optimizer Auto Scaling Group Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_AutoScalingGroupRecommendation.html

The primary key for this table is **auto_scaling_group_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|auto_scaling_group_arn (PK)|`utf8`|
|auto_scaling_group_name|`utf8`|
|current_configuration|`json`|
|current_performance_risk|`utf8`|
|effective_recommendation_preferences|`json`|
|finding|`utf8`|
|inferred_workload_types|`list<item: utf8, nullable>`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|look_back_period_in_days|`float64`|
|recommendation_options|`json`|
|utilization_metrics|`json`|