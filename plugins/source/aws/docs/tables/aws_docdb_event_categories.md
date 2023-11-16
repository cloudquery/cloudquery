# Table: aws_docdb_event_categories

This table shows data for Amazon DocumentDB Event Categories.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventCategoriesMap.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|event_categories|`list<item: utf8, nullable>`|
|source_type|`utf8`|