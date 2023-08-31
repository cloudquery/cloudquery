# Table: aws_autoscaling_group_scaling_policies

This table shows data for Auto Scaling Group Scaling Policies.

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ScalingPolicy.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_autoscaling_groups](aws_autoscaling_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|group_arn|`utf8`|
|arn (PK)|`utf8`|
|adjustment_type|`utf8`|
|alarms|`json`|
|auto_scaling_group_name|`utf8`|
|cooldown|`int64`|
|enabled|`bool`|
|estimated_instance_warmup|`int64`|
|metric_aggregation_type|`utf8`|
|min_adjustment_magnitude|`int64`|
|min_adjustment_step|`int64`|
|policy_arn|`utf8`|
|policy_name|`utf8`|
|policy_type|`utf8`|
|predictive_scaling_configuration|`json`|
|scaling_adjustment|`int64`|
|step_adjustments|`json`|
|target_tracking_configuration|`json`|