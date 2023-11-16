# Table: aws_eventbridge_connections

This table shows data for Amazon EventBridge Connections.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Connection.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|authorization_type|`utf8`|
|connection_arn|`utf8`|
|connection_state|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|last_authorized_time|`timestamp[us, tz=UTC]`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|state_reason|`utf8`|