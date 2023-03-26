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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|application_source|JSON|
|scaling_instructions|JSON|
|scaling_plan_name (PK)|String|
|scaling_plan_version|Int|
|status_code|String|
|creation_time|Timestamp|
|status_message|String|
|status_start_time|Timestamp|