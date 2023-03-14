# Table: heroku_team_members

This table shows data for Heroku Team Members.

https://devcenter.heroku.com/articles/platform-api-reference#team-member

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
|email|String|
|federated|Bool|
|identity_provider|JSON|
|role|String|
|two_factor_authentication|Bool|
|updated_at|Timestamp|
|user|JSON|