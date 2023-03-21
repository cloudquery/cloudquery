# Table: heroku_team_invitations

This table shows data for Heroku Team Invitations.

https://devcenter.heroku.com/articles/platform-api-reference#team-invitation

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
|invited_by|JSON|
|role|String|
|team|JSON|
|updated_at|Timestamp|
|user|JSON|