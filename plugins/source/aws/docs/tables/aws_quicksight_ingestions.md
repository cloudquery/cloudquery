# Table: aws_quicksight_ingestions



The composite primary key for this table is (**arn**, **data_set_arn**).

## Relations
This table depends on [aws_quicksight_data_sets](aws_quicksight_data_sets.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|data_set_arn (PK)|String|
|created_time|Timestamp|
|ingestion_status|String|
|error_info|JSON|
|ingestion_id|String|
|ingestion_size_in_bytes|Int|
|ingestion_time_in_seconds|Int|
|queue_info|JSON|
|request_source|String|
|request_type|String|
|row_info|JSON|