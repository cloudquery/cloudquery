# Table: aws_applicationautoscaling_scaling_activities

This table shows data for Applicationautoscaling Scaling Activities.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingActivity.html

The composite primary key for this table is (**account_id**, **region**, **resource_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|activity_id|String|
|cause|String|
|description|String|
|resource_id (PK)|String|
|scalable_dimension|String|
|service_namespace|String|
|start_time|Timestamp|
|status_code|String|
|details|String|
|end_time|Timestamp|
|not_scaled_reasons|JSON|
|status_message|String|