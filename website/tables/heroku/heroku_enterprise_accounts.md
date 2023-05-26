# Table: heroku_enterprise_accounts

This table shows data for Heroku Enterprise Accounts.

https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account

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
|identity_provider|json|
|name|utf8|
|permissions|list<item: utf8, nullable>|
|trial|bool|
|updated_at|timestamp[us, tz=UTC]|