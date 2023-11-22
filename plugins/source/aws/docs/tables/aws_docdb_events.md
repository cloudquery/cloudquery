# Table: aws_docdb_events

This table shows data for Amazon DocumentDB Events.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html

The composite primary key for this table is (**account_id**, **region**, **date**, **source_arn**, **source_identifier**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|date (PK)|`timestamp[us, tz=UTC]`|
|event_categories|`list<item: utf8, nullable>`|
|message|`utf8`|
|source_arn (PK)|`utf8`|
|source_identifier (PK)|`utf8`|
|source_type|`utf8`|