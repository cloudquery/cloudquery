# Table: aws_timestream_tables

This table shows data for Timestream Tables.

https://docs.aws.amazon.com/timestream/latest/developerguide/API_Table.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_timestream_databases](aws_timestream_databases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|database_name|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|magnetic_store_write_properties|`json`|
|retention_properties|`json`|
|schema|`json`|
|table_name|`utf8`|
|table_status|`utf8`|