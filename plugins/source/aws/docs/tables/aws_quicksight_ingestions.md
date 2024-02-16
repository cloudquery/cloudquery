# Table: aws_quicksight_ingestions

This table shows data for QuickSight Ingestions.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Ingestion.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **data_set_arn**, **arn**).
## Relations

This table depends on [aws_quicksight_data_sets](aws_quicksight_data_sets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|data_set_arn|`utf8`|
|arn|`utf8`|
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