# Table: aws_lightsail_database_log_events

This table shows data for Lightsail Database Log Events.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRelationalDatabaseLogEvents.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_databases](aws_lightsail_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|database_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|log_stream_name|`utf8`|