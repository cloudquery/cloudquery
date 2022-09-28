# Table: heroku_oauth_clients
https://devcenter.heroku.com/articles/platform-api-reference#o-auth-client-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|created_at|Timestamp|
|id (PK)|String|
|ignores_delinquent|Bool|
|name|String|
|redirect_uri|String|
|secret|String|
|updated_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|