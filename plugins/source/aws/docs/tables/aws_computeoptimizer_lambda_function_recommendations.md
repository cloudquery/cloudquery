# Table: aws_computeoptimizer_lambda_function_recommendations

This table shows data for Compute Optimizer AWS Lambda Function Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_LambdaFunctionRecommendation.html

The primary key for this table is **function_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|current_memory_size|`int64`|
|current_performance_risk|`utf8`|
|finding|`utf8`|
|finding_reason_codes|`list<item: utf8, nullable>`|
|function_arn (PK)|`utf8`|
|function_version|`utf8`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|lookback_period_in_days|`float64`|
|memory_size_recommendation_options|`json`|
|number_of_invocations|`int64`|
|utilization_metrics|`json`|