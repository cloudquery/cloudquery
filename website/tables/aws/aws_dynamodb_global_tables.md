# Table: aws_dynamodb_global_tables

This table shows data for Amazon DynamoDB Global Tables.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalTableDescription.html
This table only contains version 2017.11.29 (Legacy) Global Tables. See aws_dynamodb_tables for version 2019.11.21 (Current) Global Tables.

The composite primary key for this table is (**region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|global_table_arn|`utf8`|
|global_table_name|`utf8`|
|global_table_status|`utf8`|
|replication_group|`json`|