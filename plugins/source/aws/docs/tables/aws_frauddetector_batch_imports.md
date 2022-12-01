# Table: aws_frauddetector_batch_imports

https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchImport.html

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
|completion_time|String|
|event_type_name|String|
|failed_records_count|Int|
|failure_reason|String|
|iam_role_arn|String|
|input_path|String|
|job_id|String|
|output_path|String|
|processed_records_count|Int|
|start_time|String|
|status|String|
|total_records_count|Int|