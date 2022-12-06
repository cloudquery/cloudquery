# Table: heroku_team_invitations

https://devcenter.heroku.com/articles/platform-api-reference#team-invitation

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|created_at|Timestamp|
|id (PK)|String|
|invited_by|JSON|
|role|String|
|team|JSON|
|updated_at|Timestamp|
|user|JSON|