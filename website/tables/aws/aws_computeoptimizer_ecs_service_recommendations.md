# Table: aws_computeoptimizer_ecs_service_recommendations

This table shows data for Compute Optimizer Amazon Elastic Container Service (ECS) Service Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_ECSServiceRecommendation.html

The primary key for this table is **service_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|current_performance_risk|String|
|current_service_configuration|JSON|
|finding|String|
|finding_reason_codes|StringArray|
|last_refresh_timestamp|Timestamp|
|launch_type|String|
|lookback_period_in_days|Float|
|service_arn (PK)|String|
|service_recommendation_options|JSON|
|utilization_metrics|JSON|