# Table: aws_computeoptimizer_ecs_service_recommendations

This table shows data for Compute Optimizer Amazon Elastic Container Service (ECS) Service Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_ECSServiceRecommendation.html

The primary key for this table is **service_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|current_performance_risk|`utf8`|
|current_service_configuration|`json`|
|finding|`utf8`|
|finding_reason_codes|`list<item: utf8, nullable>`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|launch_type|`utf8`|
|lookback_period_in_days|`float64`|
|service_arn (PK)|`utf8`|
|service_recommendation_options|`json`|
|utilization_metrics|`json`|