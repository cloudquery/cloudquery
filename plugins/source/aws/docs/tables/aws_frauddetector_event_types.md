# Table: aws_frauddetector_event_types

https://docs.aws.amazon.com/frauddetector/latest/api/API_EventType.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|created_time|String|
|description|String|
|entity_types|StringArray|
|event_ingestion|String|
|event_variables|StringArray|
|ingested_event_statistics|JSON|
|labels|StringArray|
|last_updated_time|String|
|name|String|