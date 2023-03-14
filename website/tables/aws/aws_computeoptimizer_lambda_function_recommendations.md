# Table: aws_computeoptimizer_lambda_function_recommendations

This table shows data for Compute Optimizer AWS Lambda Function Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_LambdaFunctionRecommendation.html

The primary key for this table is **function_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|current_memory_size|Int|
|current_performance_risk|String|
|finding|String|
|finding_reason_codes|StringArray|
|function_arn (PK)|String|
|function_version|String|
|last_refresh_timestamp|Timestamp|
|lookback_period_in_days|Float|
|memory_size_recommendation_options|JSON|
|number_of_invocations|Int|
|utilization_metrics|JSON|