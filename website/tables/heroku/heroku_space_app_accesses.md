# Table: heroku_space_app_accesses

This table shows data for Heroku Space App Accesses.

https://devcenter.heroku.com/articles/platform-api-reference#space-app-access

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|permissions|JSON|
|space|JSON|
|updated_at|Timestamp|
|user|JSON|