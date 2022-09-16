
# Table: aws_glue_triggers
Information about a specific trigger
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the trigger.|
|tags|jsonb|Resource tags.|
|description|text|A description of this trigger|
|event_batching_condition_size|bigint|Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires|
|event_batching_condition_window|bigint|Window of time in seconds after which EventBridge event trigger fires|
|id|text|Reserved for future use|
|name|text|The name of the trigger|
|predicate_logical|text|An optional field if only one condition is listed|
|schedule|text|A cron expression used to specify the schedule (see Time-Based Schedules for Jobs and Crawlers (https://docsawsamazoncom/glue/latest/dg/monitor-data-warehouse-schedulehtml) For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *)|
|state|text|The current state of the trigger|
|type|text|The type of trigger that this is|
|workflow_name|text|The name of the workflow associated with the trigger|
