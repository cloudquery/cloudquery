# Table: aws_elasticache_events

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Event.html

The composite primary key for this table is (**date**, **source_identifier**, **source_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|date (PK)|Timestamp|
|message|String|
|source_identifier (PK)|String|
|source_type (PK)|String|