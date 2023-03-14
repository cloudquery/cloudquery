# Table: heroku_enterprise_account_members

This table shows data for Heroku Enterprise Account Members.

https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account-member

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|enterprise_account|JSON|
|identity_provider|JSON|
|permissions|JSON|
|two_factor_authentication|Bool|
|user|JSON|