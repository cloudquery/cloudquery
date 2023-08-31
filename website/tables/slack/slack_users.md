# Table: slack_users

This table shows data for Slack Users.

https://api.slack.com/methods/users.list

The primary key for this table is **id**.

## Relations

The following tables depend on slack_users:
  - [slack_user_presences](slack_user_presences)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|updated|`timestamp[us, tz=UTC]`|
|team_id|`utf8`|
|name|`utf8`|
|deleted|`bool`|
|color|`utf8`|
|real_name|`utf8`|
|tz|`utf8`|
|tz_label|`utf8`|
|tz_offset|`int64`|
|profile|`json`|
|is_bot|`bool`|
|is_admin|`bool`|
|is_owner|`bool`|
|is_primary_owner|`bool`|
|is_restricted|`bool`|
|is_ultra_restricted|`bool`|
|is_stranger|`bool`|
|is_app_user|`bool`|
|is_invited_user|`bool`|
|has_files|`bool`|
|locale|`utf8`|
|enterprise_user|`json`|