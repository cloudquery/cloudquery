# Table: heroku_team_app_permissions

This table shows data for Heroku Team App Permissions.

https://devcenter.heroku.com/articles/platform-api-reference#team-app-permission

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id (PK)|uuid|
|_cq_parent_id|uuid|
|description|utf8|
|name|utf8|