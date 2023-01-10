# Table: aws_applicationautoscaling_policies

https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPolicy.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|creation_time|Timestamp|
|policy_arn|String|
|policy_name|String|
|policy_type|String|
|resource_id|String|
|scalable_dimension|String|
|service_namespace|String|
|alarms|JSON|
|step_scaling_policy_configuration|JSON|
|target_tracking_scaling_policy_configuration|JSON|