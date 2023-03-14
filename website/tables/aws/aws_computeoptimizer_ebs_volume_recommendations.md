# Table: aws_computeoptimizer_ebs_volume_recommendations

This table shows data for Compute Optimizer Amazon Elastic Block Store (EBS) Volume Recommendations.

https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_VolumeRecommendation.html

The primary key for this table is **volume_arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|current_configuration|JSON|
|current_performance_risk|String|
|finding|String|
|last_refresh_timestamp|Timestamp|
|look_back_period_in_days|Float|
|utilization_metrics|JSON|
|volume_arn (PK)|String|
|volume_recommendation_options|JSON|