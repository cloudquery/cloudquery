# Table: aws_elasticache_events

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Event.html

The primary key for this table is **_event_hash**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|_event_hash (PK)|String|
|date|Timestamp|
|message|String|
|source_identifier|String|
|source_type|String|