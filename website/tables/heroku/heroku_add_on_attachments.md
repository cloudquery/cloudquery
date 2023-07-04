# Table: heroku_add_on_attachments

This table shows data for Heroku Add On Attachments.

https://devcenter.heroku.com/articles/platform-api-reference#add-on-attachment

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|addon|`json`|
|app|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|log_input_url|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|web_url|`utf8`|