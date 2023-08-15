# Table: aws_timestream_databases

This table shows data for Timestream Databases.

https://docs.aws.amazon.com/timestream/latest/developerguide/API_Database.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_timestream_databases:
  - [aws_timestream_tables](aws_timestream_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|database_name|`utf8`|
|kms_key_id|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|table_count|`int64`|