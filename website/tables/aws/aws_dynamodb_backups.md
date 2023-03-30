# Table: aws_dynamodb_backups

This table shows data for Amazon DynamoDB Backups.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BackupDescription.html

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
|backup_details|JSON|
|source_table_details|JSON|
|source_table_feature_details|JSON|