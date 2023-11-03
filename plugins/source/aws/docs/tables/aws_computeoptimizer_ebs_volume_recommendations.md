# Table: aws_computeoptimizer_ebs_volume_recommendations

This table shows data for Compute Optimizer Amazon Elastic Block Store (EBS) Volume Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_VolumeRecommendation.html

The primary key for this table is **volume_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|tags|`json`|
|current_configuration|`json`|
|current_performance_risk|`utf8`|
|finding|`utf8`|
|last_refresh_timestamp|`timestamp[us, tz=UTC]`|
|look_back_period_in_days|`float64`|
|utilization_metrics|`json`|
|volume_arn (PK)|`utf8`|
|volume_recommendation_options|`json`|