# Table: hackernews_items

https://github.com/HackerNews/API#items

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|deleted|Bool|
|type|String|
|by|String|
|time|Int|
|text|String|
|dead|Bool|
|parent|Int|
|kids|IntArray|
|url|String|
|score|Int|
|title|String|
|parts|IntArray|
|descendants|Int|