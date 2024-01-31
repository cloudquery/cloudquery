# Table: aws_applicationautoscaling_policies

This table shows data for Application Auto Scaling Policies.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingPolicy.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|policy_arn|`utf8`|
|policy_name|`utf8`|
|policy_type|`utf8`|
|resource_id|`utf8`|
|scalable_dimension|`utf8`|
|service_namespace|`utf8`|
|alarms|`json`|
|step_scaling_policy_configuration|`json`|
|target_tracking_scaling_policy_configuration|`json`|