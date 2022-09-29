# Table: heroku_oauth_authorizations
https://devcenter.heroku.com/articles/platform-api-reference#o-auth-authorization-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|access_token|JSON|
|client|JSON|
|created_at|Timestamp|
|grant|JSON|
|id (PK)|String|
|refresh_token|JSON|
|scope|StringArray|
|updated_at|Timestamp|
|user|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|