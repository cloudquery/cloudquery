# Table: aws_dynamodb_table_continuous_backups


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_dynamodb_tables`](aws_dynamodb_tables.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|table_arn|String|
|continuous_backups_status|String|
|point_in_time_recovery_description|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|