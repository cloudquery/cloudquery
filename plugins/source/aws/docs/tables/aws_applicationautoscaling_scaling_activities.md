# Table: aws_applicationautoscaling_scaling_activities

This table shows data for Application Auto Scaling Scaling Activities.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingActivity.html

The composite primary key for this table is (**account_id**, **region**, **resource_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|activity_id|`utf8`|
|cause|`utf8`|
|description|`utf8`|
|resource_id (PK)|`utf8`|
|scalable_dimension|`utf8`|
|service_namespace|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|status_code|`utf8`|
|details|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|not_scaled_reasons|`json`|
|status_message|`utf8`|