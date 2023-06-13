# Table: heroku_permission_entities

This table shows data for Heroku Permission Entities.

https://devcenter.heroku.com/articles/platform-api-reference#permission-entity

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|team_id|`utf8`|
|type|`utf8`|
|users|`json`|