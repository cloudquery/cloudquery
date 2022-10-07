# Table: heroku_buildpack_installations

https://devcenter.heroku.com/articles/platform-api-reference#buildpack-installation-attributes

The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|buildpack|JSON|
|ordinal|Int|