# Table: aws_apprunner_connections

This table shows data for AWS App Runner Connections.

https://docs.aws.amazon.com/apprunner/latest/api/API_Connection.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|connection_arn|`utf8`|
|connection_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|provider_type|`utf8`|
|status|`utf8`|