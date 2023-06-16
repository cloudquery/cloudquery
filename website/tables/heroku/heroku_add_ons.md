# Table: heroku_add_ons

This table shows data for Heroku Add Ons.

https://devcenter.heroku.com/articles/platform-api-reference#add-on

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|actions|`json`|
|addon_service|`json`|
|app|`json`|
|billed_price|`json`|
|billing_entity|`json`|
|config_vars|`list<item: utf8, nullable>`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|plan|`json`|
|provider_id|`utf8`|
|state|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|web_url|`utf8`|