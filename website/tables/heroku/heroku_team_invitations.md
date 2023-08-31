# Table: heroku_team_invitations

This table shows data for Heroku Team Invitations.

https://devcenter.heroku.com/articles/platform-api-reference#team-invitation

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|invited_by|`json`|
|role|`utf8`|
|team|`json`|
|updated_at|`timestamp[us, tz=UTC]`|
|user|`json`|