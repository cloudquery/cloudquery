# Table: notion_pages_table

This table shows data for Notion Pages Table.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|object|`utf8`|
|id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|last_edited_time|`timestamp[us, tz=UTC]`|
|created_by|`json`|
|last_edited_by|`json`|
|cover|`json`|
|icon|`json`|
|parent|`json`|
|archived|`bool`|
|properties|`json`|
|url|`utf8`|
|public_url|`utf8`|