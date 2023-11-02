# Table: aws_glue_connections

This table shows data for Glue Connections.

https://docs.aws.amazon.com/glue/latest/webapi/API_Connection.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|connection_properties|`json`|
|connection_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_updated_by|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|match_criteria|`list<item: utf8, nullable>`|
|name|`utf8`|
|physical_connection_requirements|`json`|