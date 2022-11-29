# Table: aws_apprunner_operations

https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_apprunner_services](aws_apprunner_services.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|ended_at|Timestamp|
|id|String|
|started_at|Timestamp|
|status|String|
|target_arn|String|
|type|String|
|updated_at|Timestamp|