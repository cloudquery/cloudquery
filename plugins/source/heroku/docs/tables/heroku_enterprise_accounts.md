# Table: heroku_enterprise_accounts
https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|created_at|Timestamp|
|id (PK)|String|
|identity_provider|JSON|
|name|String|
|permissions|StringArray|
|trial|Bool|
|updated_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|