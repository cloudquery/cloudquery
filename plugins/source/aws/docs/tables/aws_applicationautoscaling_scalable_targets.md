# Table: aws_applicationautoscaling_scalable_targets

This table shows data for Application Auto Scaling Scalable Targets.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalableTarget.html

The composite primary key for this table is (**account_id**, **region**, **resource_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|max_capacity|`int64`|
|min_capacity|`int64`|
|resource_id (PK)|`utf8`|
|role_arn|`utf8`|
|scalable_dimension|`utf8`|
|service_namespace|`utf8`|
|scalable_target_arn|`utf8`|
|suspended_state|`json`|