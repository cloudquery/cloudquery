# Table: aws_frauddetector_batch_imports

This table shows data for Amazon Fraud Detector Batch Imports.

https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchImport.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|completion_time|`utf8`|
|event_type_name|`utf8`|
|failed_records_count|`int64`|
|failure_reason|`utf8`|
|iam_role_arn|`utf8`|
|input_path|`utf8`|
|job_id|`utf8`|
|output_path|`utf8`|
|processed_records_count|`int64`|
|start_time|`utf8`|
|status|`utf8`|
|total_records_count|`int64`|