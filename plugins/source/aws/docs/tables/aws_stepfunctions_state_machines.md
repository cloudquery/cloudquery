# Table: aws_stepfunctions_state_machines

https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html

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
|creation_date|Timestamp|
|definition|String|
|name|String|
|role_arn|String|
|type|String|
|logging_configuration|JSON|
|status|String|
|tracing_configuration|JSON|