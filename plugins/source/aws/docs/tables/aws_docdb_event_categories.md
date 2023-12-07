# Table: aws_docdb_event_categories

This table shows data for Amazon DocumentDB Event Categories.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventCategoriesMap.html

The composite primary key for this table is (**account_id**, **region**, **source_type**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|event_categories|`list<item: utf8, nullable>`|
|source_type (PK)|`utf8`|