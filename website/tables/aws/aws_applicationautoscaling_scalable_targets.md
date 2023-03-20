# Table: aws_applicationautoscaling_scalable_targets

This table shows data for Applicationautoscaling Scalable Targets.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalableTarget.html

The composite primary key for this table is (**account_id**, **region**, **resource_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|creation_time|Timestamp|
|max_capacity|Int|
|min_capacity|Int|
|resource_id (PK)|String|
|role_arn|String|
|scalable_dimension|String|
|service_namespace|String|
|suspended_state|JSON|