# Table: heroku_enterprise_accounts

https://devcenter.heroku.com/articles/platform-api-reference#enterprise-account

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
|identity_provider|JSON|
|name|String|
|permissions|StringArray|
|trial|Bool|
|updated_at|Timestamp|