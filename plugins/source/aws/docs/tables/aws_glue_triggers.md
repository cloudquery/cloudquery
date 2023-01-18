# Table: aws_glue_triggers

https://docs.aws.amazon.com/glue/latest/webapi/API_Trigger.html

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
|tags|JSON|
|actions|JSON|
|description|String|
|event_batching_condition|JSON|
|id|String|
|name|String|
|predicate|JSON|
|schedule|String|
|state|String|
|type|String|
|workflow_name|String|