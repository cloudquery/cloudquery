# Table: slack_conversation_bookmarks

This table shows data for Slack Conversation Bookmarks.

https://api.slack.com/methods/bookmarks.list

The composite primary key for this table is (**team_id**, **id**, **channel_id**).

## Relations

This table depends on [slack_conversations](slack_conversations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|team_id (PK)|utf8|
|id (PK)|utf8|
|channel_id (PK)|utf8|
|date_created|timestamp[us, tz=UTC]|
|date_updated|timestamp[us, tz=UTC]|
|title|utf8|
|link|utf8|
|emoji|utf8|
|icon_url|utf8|
|type|utf8|
|rank|utf8|
|last_updated_by_user_id|utf8|
|last_updated_by_team_id|utf8|
|shortcut_id|utf8|
|entity_id|utf8|
|app_id|utf8|