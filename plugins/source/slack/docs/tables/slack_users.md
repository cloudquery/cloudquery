# Table: slack_users

https://api.slack.com/methods/users.list

The primary key for this table is **id**.

## Relations

The following tables depend on slack_users:
  - [slack_user_presences](slack_user_presences.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|team_id|String|
|name|String|
|deleted|Bool|
|color|String|
|real_name|String|
|tz|String|
|tz_label|String|
|tz_offset|Int|
|profile|JSON|
|is_bot|Bool|
|is_admin|Bool|
|is_owner|Bool|
|is_primary_owner|Bool|
|is_restricted|Bool|
|is_ultra_restricted|Bool|
|is_stranger|Bool|
|is_app_user|Bool|
|is_invited_user|Bool|
|has_files|Bool|
|locale|String|
|updated|Timestamp|
|enterprise_user|JSON|