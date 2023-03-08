# Table: heroku_oauth_clients

https://devcenter.heroku.com/articles/platform-api-reference#o-auth-client

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
|ignores_delinquent|Bool|
|name|String|
|redirect_uri|String|
|secret|String|
|updated_at|Timestamp|