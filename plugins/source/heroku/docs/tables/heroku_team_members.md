# Table: heroku_team_members

https://devcenter.heroku.com/articles/platform-api-reference#team-member

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|created_at|Timestamp|
|email|String|
|federated|Bool|
|id (PK)|String|
|identity_provider|JSON|
|role|String|
|two_factor_authentication|Bool|
|updated_at|Timestamp|
|user|JSON|