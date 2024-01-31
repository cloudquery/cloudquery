# Table: aws_dynamodb_tables

This table shows data for Amazon DynamoDB Tables.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_dynamodb_tables:
  - [aws_dynamodb_table_continuous_backups](aws_dynamodb_table_continuous_backups.md)
  - [aws_dynamodb_table_replica_auto_scalings](aws_dynamodb_table_replica_auto_scalings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|archival_summary|`json`|
|attribute_definitions|`json`|
|billing_mode_summary|`json`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|deletion_protection_enabled|`bool`|
|global_secondary_indexes|`json`|
|global_table_version|`utf8`|
|item_count|`int64`|
|key_schema|`json`|
|latest_stream_arn|`utf8`|
|latest_stream_label|`utf8`|
|local_secondary_indexes|`json`|
|provisioned_throughput|`json`|
|replicas|`json`|
|restore_summary|`json`|
|sse_description|`json`|
|stream_specification|`json`|
|table_arn|`utf8`|
|table_class_summary|`json`|
|table_id|`utf8`|
|table_name|`utf8`|
|table_size_bytes|`int64`|
|table_status|`utf8`|