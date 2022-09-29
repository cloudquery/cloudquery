# Table: heroku_team_members
https://devcenter.heroku.com/articles/platform-api-reference#team-member-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|created_at|Timestamp|
|email|String|
|federated|Bool|
|id (PK)|String|
|identity_provider|JSON|
|role|String|
|two_factor_authentication|Bool|
|updated_at|Timestamp|
|user|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|