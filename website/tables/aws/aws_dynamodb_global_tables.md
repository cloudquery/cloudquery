# Table: aws_dynamodb_global_tables

This table shows data for Amazon DynamoDB Global Tables.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalTableDescription.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|arn (PK)|String|
|tags|JSON|
|creation_date_time|Timestamp|
|global_table_arn|String|
|global_table_name|String|
|global_table_status|String|
|replication_group|JSON|