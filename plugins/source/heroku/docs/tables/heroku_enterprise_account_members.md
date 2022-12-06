# Table: heroku_enterprise_account_members

https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account-member

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|enterprise_account|JSON|
|id (PK)|String|
|identity_provider|JSON|
|permissions|JSON|
|two_factor_authentication|Bool|
|user|JSON|