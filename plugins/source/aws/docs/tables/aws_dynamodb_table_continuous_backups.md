# Table: aws_dynamodb_table_continuous_backups

This table shows data for Amazon DynamoDB Table Continuous Backups.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ContinuousBackupsDescription.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **table_arn**.
## Relations

This table depends on [aws_dynamodb_tables](aws_dynamodb_tables.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|table_arn|`utf8`|
|continuous_backups_status|`utf8`|
|point_in_time_recovery_description|`json`|