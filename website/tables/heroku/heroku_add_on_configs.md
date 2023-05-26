# Table: heroku_add_on_configs

This table shows data for Heroku Add On Configs.

https://devcenter.heroku.com/articles/platform-api-reference#add-on-config

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|name|utf8|
|value|utf8|