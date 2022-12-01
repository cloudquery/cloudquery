# Table: aws_dynamodb_table_continuous_backups

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ContinuousBackupsDescription.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_dynamodb_tables](aws_dynamodb_tables.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|table_arn|String|
|continuous_backups_status|String|
|point_in_time_recovery_description|JSON|