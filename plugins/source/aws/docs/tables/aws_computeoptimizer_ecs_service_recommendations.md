# Table: aws_computeoptimizer_ecs_service_recommendations

This table shows data for Compute Optimizer Amazon Elastic Container Service (ECS) Service Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_ECSServiceRecommendation.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **service_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|current_performance_risk|`utf8`|
|current_service_configuration|`json`|
|effective_recommendation_preferences|`json`|
|finding|`utf8`|
|finding_reason_codes|`list<item: utf8, nullable>`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|launch_type|`utf8`|
|lookback_period_in_days|`float64`|
|service_arn|`utf8`|
|service_recommendation_options|`json`|
|utilization_metrics|`json`|