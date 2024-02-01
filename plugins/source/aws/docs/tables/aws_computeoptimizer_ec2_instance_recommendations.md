# Table: aws_computeoptimizer_ec2_instance_recommendations

This table shows data for Compute Optimizer Amazon Elastic Compute Cloud (EC2) Instance Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_InstanceRecommendation.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **instance_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|current_instance_gpu_info|`json`|
|current_instance_type|`utf8`|
|current_performance_risk|`utf8`|
|effective_recommendation_preferences|`json`|
|external_metric_status|`json`|
|finding|`utf8`|
|finding_reason_codes|`list<item: utf8, nullable>`|
|idle|`utf8`|
|inferred_workload_types|`list<item: utf8, nullable>`|
|instance_arn|`utf8`|
|instance_name|`utf8`|
|instance_state|`utf8`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|look_back_period_in_days|`float64`|
|recommendation_options|`json`|
|recommendation_sources|`json`|
|utilization_metrics|`json`|