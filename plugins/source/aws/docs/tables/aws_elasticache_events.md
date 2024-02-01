# Table: aws_elasticache_events

This table shows data for Elasticache Events.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Event.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **_event_hash**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|_event_hash|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|source_identifier|`utf8`|
|source_type|`utf8`|