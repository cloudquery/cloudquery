# Table: heroku_buildpack_installations

This table shows data for Heroku Buildpack Installations.

https://devcenter.heroku.com/articles/platform-api-reference#buildpack-installation

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|buildpack|json|
|ordinal|int64|