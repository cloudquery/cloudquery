# Table: heroku_oauth_authorizations

https://devcenter.heroku.com/articles/platform-api-reference#o-auth-authorization-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|access_token|JSON|
|client|JSON|
|created_at|Timestamp|
|grant|JSON|
|id (PK)|String|
|refresh_token|JSON|
|scope|StringArray|
|updated_at|Timestamp|
|user|JSON|