# Table: heroku_oauth_clients

This table shows data for Heroku OAuth Clients.

https://devcenter.heroku.com/articles/platform-api-reference#o-auth-client

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|created_at|timestamp[us, tz=UTC]|
|ignores_delinquent|bool|
|name|utf8|
|redirect_uri|utf8|
|secret|utf8|
|updated_at|timestamp[us, tz=UTC]|