# Table: aws_autoscaling_plan_resources

This table shows data for Auto Scaling Plan Resources.

https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPlanResource.html

The composite primary key for this table is (**account_id**, **region**, **resource_id**, **scaling_plan_name**).

## Relations

This table depends on [aws_autoscaling_plans](aws_autoscaling_plans).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|resource_id (PK)|`utf8`|
|scalable_dimension|`utf8`|
|scaling_plan_name (PK)|`utf8`|
|scaling_plan_version|`int64`|
|scaling_status_code|`utf8`|
|service_namespace|`utf8`|
|scaling_policies|`json`|
|scaling_status_message|`utf8`|