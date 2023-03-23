# Table: aws_applicationautoscaling_scheduled_actions

This table shows data for Application Auto Scaling Scheduled Actions.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScheduledAction.html

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
|resource_id|String|
|schedule|String|
|scheduled_action_arn|String|
|scheduled_action_name|String|
|service_namespace|String|
|end_time|Timestamp|
|scalable_dimension|String|
|scalable_target_action|JSON|
|start_time|Timestamp|
|timezone|String|