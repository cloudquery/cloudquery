# Table: aws_elasticache_events

This table shows data for Elasticache Events.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Event.html

The primary key for this table is **_event_hash**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|_event_hash (PK)|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|message|`utf8`|
|source_identifier|`utf8`|
|source_type|`utf8`|