# Table: slack_user_presences

This table shows data for Slack User Presences.

https://api.slack.com/methods/users.getPresence

The primary key for this table is **user_id**.

## Relations

This table depends on [slack_users](slack_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|user_id (PK)|`utf8`|
|last_activity|`timestamp[us, tz=UTC]`|
|presence|`utf8`|
|online|`bool`|
|auto_away|`bool`|
|manual_away|`bool`|
|connection_count|`int64`|