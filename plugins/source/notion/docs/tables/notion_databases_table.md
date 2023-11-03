# Table: notion_databases_table

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|object|`utf8`|
|id (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|last_edited_time|`timestamp[us, tz=UTC]`|
|created_by|`json`|
|last_edited_by|`json`|
|title|`json`|
|description|`json`|
|is_inline|`bool`|
|cover|`json`|
|icon|`json`|
|parent|`json`|
|archived|`bool`|
|properties|`json`|
|url|`utf8`|
|public_url|`utf8`|