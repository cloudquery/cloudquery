# Table: aws_applicationautoscaling_scalable_targets

This table shows data for Application Auto Scaling Scalable Targets.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalableTarget.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **resource_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|max_capacity|`int64`|
|min_capacity|`int64`|
|resource_id|`utf8`|
|role_arn|`utf8`|
|scalable_dimension|`utf8`|
|service_namespace|`utf8`|
|scalable_target_arn|`utf8`|
|suspended_state|`json`|