# Table: aws_autoscaling_plan_resources

This table shows data for Auto Scaling Plan Resources.

https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPlanResource.html

The composite primary key for this table is (**account_id**, **region**, **resource_id**, **scaling_plan_name**).

## Relations

This table depends on [aws_autoscaling_plans](aws_autoscaling_plans).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|resource_id (PK)|String|
|scalable_dimension|String|
|scaling_plan_name (PK)|String|
|scaling_plan_version|Int|
|scaling_status_code|String|
|service_namespace|String|
|scaling_policies|JSON|
|scaling_status_message|String|