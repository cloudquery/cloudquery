# Table: aws_dynamodb_exports

This table shows data for Amazon DynamoDB Exports.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ExportDescription.html

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
|billed_size_bytes|Int|
|client_token|String|
|end_time|Timestamp|
|export_arn|String|
|export_format|String|
|export_manifest|String|
|export_status|String|
|export_time|Timestamp|
|failure_code|String|
|failure_message|String|
|item_count|Int|
|s3_bucket|String|
|s3_bucket_owner|String|
|s3_prefix|String|
|s3_sse_algorithm|String|
|s3_sse_kms_key_id|String|
|start_time|Timestamp|
|table_arn|String|
|table_id|String|