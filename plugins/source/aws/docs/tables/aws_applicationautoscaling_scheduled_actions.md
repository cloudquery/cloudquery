# Table: aws_applicationautoscaling_scheduled_actions

This table shows data for Application Auto Scaling Scheduled Actions.

https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScheduledAction.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|resource_id|`utf8`|
|schedule|`utf8`|
|scheduled_action_arn|`utf8`|
|scheduled_action_name|`utf8`|
|service_namespace|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|scalable_dimension|`utf8`|
|scalable_target_action|`json`|
|start_time|`timestamp[us, tz=UTC]`|
|timezone|`utf8`|