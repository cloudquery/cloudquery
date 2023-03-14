# Table: aws_computeoptimizer_autoscaling_group_recommendations

This table shows data for Compute Optimizer Auto Scaling Group Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_AutoScalingGroupRecommendation.html

The primary key for this table is **auto_scaling_group_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|auto_scaling_group_arn (PK)|String|
|auto_scaling_group_name|String|
|current_configuration|JSON|
|current_performance_risk|String|
|effective_recommendation_preferences|JSON|
|finding|String|
|inferred_workload_types|StringArray|
|last_refresh_timestamp|Timestamp|
|look_back_period_in_days|Float|
|recommendation_options|JSON|
|utilization_metrics|JSON|