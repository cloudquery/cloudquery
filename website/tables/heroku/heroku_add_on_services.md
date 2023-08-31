# Table: heroku_add_on_services

This table shows data for Heroku Add On Services.

https://devcenter.heroku.com/articles/platform-api-reference#add-on-service

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|cli_plugin_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|human_name|`utf8`|
|name|`utf8`|
|state|`utf8`|
|supports_multiple_installations|`bool`|
|supports_sharing|`bool`|
|updated_at|`timestamp[us, tz=UTC]`|