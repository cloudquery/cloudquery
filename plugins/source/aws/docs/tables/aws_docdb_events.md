# Table: aws_docdb_events

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|date|Timestamp|
|event_categories|StringArray|
|message|String|
|source_arn|String|
|source_identifier|String|
|source_type|String|