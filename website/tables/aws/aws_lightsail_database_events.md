# Table: aws_lightsail_database_events

This table shows data for Lightsail Database Events.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseEvent.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lightsail_databases](aws_lightsail_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|database_arn|String|
|created_at|Timestamp|
|event_categories|StringArray|
|message|String|
|resource|String|