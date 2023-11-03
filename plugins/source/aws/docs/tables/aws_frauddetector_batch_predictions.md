# Table: aws_frauddetector_batch_predictions

This table shows data for Amazon Fraud Detector Batch Predictions.

https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchPrediction.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|completion_time|`utf8`|
|detector_name|`utf8`|
|detector_version|`utf8`|
|event_type_name|`utf8`|
|failure_reason|`utf8`|
|iam_role_arn|`utf8`|
|input_path|`utf8`|
|job_id|`utf8`|
|last_heartbeat_time|`utf8`|
|output_path|`utf8`|
|processed_records_count|`int64`|
|start_time|`utf8`|
|status|`utf8`|
|total_records_count|`int64`|