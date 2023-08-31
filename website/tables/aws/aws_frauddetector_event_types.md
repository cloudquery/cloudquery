# Table: aws_frauddetector_event_types

This table shows data for Amazon Fraud Detector Event Types.

https://docs.aws.amazon.com/frauddetector/latest/api/API_EventType.html

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
|created_time|`utf8`|
|description|`utf8`|
|entity_types|`list<item: utf8, nullable>`|
|event_ingestion|`utf8`|
|event_orchestration|`json`|
|event_variables|`list<item: utf8, nullable>`|
|ingested_event_statistics|`json`|
|labels|`list<item: utf8, nullable>`|
|last_updated_time|`utf8`|
|name|`utf8`|