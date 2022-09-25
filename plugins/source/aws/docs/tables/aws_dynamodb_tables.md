# Table: aws_dynamodb_tables


The primary key for this table is **arn**.

## Relations
The following tables depend on `aws_dynamodb_tables`:
  - [`aws_dynamodb_table_replica_auto_scalings`](aws_dynamodb_table_replica_auto_scalings.md)
  - [`aws_dynamodb_table_continuous_backups`](aws_dynamodb_table_continuous_backups.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|archival_summary|JSON|
|attribute_definitions|JSON|
|billing_mode_summary|JSON|
|creation_date_time|Timestamp|
|global_secondary_indexes|JSON|
|global_table_version|String|
|item_count|Int|
|key_schema|JSON|
|latest_stream_arn|String|
|latest_stream_label|String|
|local_secondary_indexes|JSON|
|provisioned_throughput|JSON|
|replicas|JSON|
|restore_summary|JSON|
|sse_description|JSON|
|stream_specification|JSON|
|table_class_summary|JSON|
|table_id|String|
|table_name|String|
|table_size_bytes|Int|
|table_status|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|