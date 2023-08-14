# Table: aws_autoscaling_plans

This table shows data for Auto Scaling Plans.

https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPlan.html

The composite primary key for this table is (**account_id**, **region**, **scaling_plan_name**).

## Relations

The following tables depend on aws_autoscaling_plans:
  - [aws_autoscaling_plan_resources](aws_autoscaling_plan_resources)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|application_source|`json`|
|scaling_instructions|`json`|
|scaling_plan_name (PK)|`utf8`|
|scaling_plan_version|`int64`|
|status_code|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|status_message|`utf8`|
|status_start_time|`timestamp[us, tz=UTC]`|