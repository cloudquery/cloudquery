# Table: aws_docdb_event_categories

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventCategoriesMap.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|event_categories|StringArray|
|source_type|String|