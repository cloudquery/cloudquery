# Table: heroku_enterprise_account_members

This table shows data for Heroku Enterprise Account Members.

https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account-member

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|enterprise_account|json|
|identity_provider|json|
|permissions|json|
|two_factor_authentication|bool|
|user|json|