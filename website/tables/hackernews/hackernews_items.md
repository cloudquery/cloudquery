# Table: hackernews_items

This table shows data for Hackernews Items.

https://github.com/HackerNews/API#items

The primary key for this table is **id**.
It supports incremental syncs based on the **id** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK) (Incremental Key)|`int64`|
|deleted|`bool`|
|type|`utf8`|
|by|`utf8`|
|time|`timestamp[us, tz=UTC]`|
|text|`utf8`|
|dead|`bool`|
|parent|`int64`|
|kids|`list<item: int64, nullable>`|
|url|`utf8`|
|score|`int64`|
|title|`utf8`|
|parts|`list<item: int64, nullable>`|
|descendants|`int64`|