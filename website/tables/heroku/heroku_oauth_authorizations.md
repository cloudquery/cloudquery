# Table: heroku_oauth_authorizations

This table shows data for Heroku OAuth Authorizations.

https://devcenter.heroku.com/articles/platform-api-reference#o-auth-authorization

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|access_token|json|
|client|json|
|created_at|timestamp[us, tz=UTC]|
|grant|json|
|refresh_token|json|
|scope|list<item: utf8, nullable>|
|updated_at|timestamp[us, tz=UTC]|
|user|json|