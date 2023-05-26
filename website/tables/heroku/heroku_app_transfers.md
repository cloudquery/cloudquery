# Table: heroku_app_transfers

This table shows data for Heroku App Transfers.

https://devcenter.heroku.com/articles/platform-api-reference#app-transfer

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|app|json|
|created_at|timestamp[us, tz=UTC]|
|owner|json|
|recipient|json|
|state|utf8|
|updated_at|timestamp[us, tz=UTC]|