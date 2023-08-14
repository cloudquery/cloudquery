# Table: aws_quicksight_ingestions

This table shows data for QuickSight Ingestions.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Ingestion.html

The composite primary key for this table is (**account_id**, **region**, **data_set_arn**, **arn**).

## Relations

This table depends on [aws_quicksight_data_sets](aws_quicksight_data_sets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|data_set_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|ingestion_status|`utf8`|
|error_info|`json`|
|ingestion_id|`utf8`|
|ingestion_size_in_bytes|`int64`|
|ingestion_time_in_seconds|`int64`|
|queue_info|`json`|
|request_source|`utf8`|
|request_type|`utf8`|
|row_info|`json`|