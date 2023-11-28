# Table: aws_dynamodb_backups

This table shows data for Amazon DynamoDB Backups.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BackupDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|backup_details|`json`|
|source_table_details|`json`|
|source_table_feature_details|`json`|