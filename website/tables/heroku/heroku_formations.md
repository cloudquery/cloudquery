# Table: heroku_formations

This table shows data for Heroku Formations.

https://devcenter.heroku.com/articles/platform-api-reference#formation

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|app|`json`|
|command|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|quantity|`int64`|
|size|`utf8`|
|type|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|