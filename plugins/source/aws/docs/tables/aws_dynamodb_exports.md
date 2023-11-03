# Table: aws_dynamodb_exports

This table shows data for Amazon DynamoDB Exports.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ExportDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|billed_size_bytes|`int64`|
|client_token|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|export_arn|`utf8`|
|export_format|`utf8`|
|export_manifest|`utf8`|
|export_status|`utf8`|
|export_time|`timestamp[us, tz=UTC]`|
|failure_code|`utf8`|
|failure_message|`utf8`|
|item_count|`int64`|
|s3_bucket|`utf8`|
|s3_bucket_owner|`utf8`|
|s3_prefix|`utf8`|
|s3_sse_algorithm|`utf8`|
|s3_sse_kms_key_id|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|table_arn|`utf8`|
|table_id|`utf8`|