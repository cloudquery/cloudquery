# Table: aws_computeoptimizer_ec2_instance_recommendations

This table shows data for Compute Optimizer Amazon Elastic Compute Cloud (EC2) Instance Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_InstanceRecommendation.html

The primary key for this table is **instance_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|current_instance_type|String|
|current_performance_risk|String|
|effective_recommendation_preferences|JSON|
|finding|String|
|finding_reason_codes|StringArray|
|inferred_workload_types|StringArray|
|instance_arn (PK)|String|
|instance_name|String|
|last_refresh_timestamp|Timestamp|
|look_back_period_in_days|Float|
|recommendation_options|JSON|
|recommendation_sources|JSON|
|utilization_metrics|JSON|