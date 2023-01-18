# Table: slack_conversation_bookmarks

https://api.slack.com/methods/bookmarks.list

The composite primary key for this table is (**team_id**, **id**, **channel_id**).

## Relations

This table depends on [slack_conversations](slack_conversations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|team_id (PK)|String|
|id (PK)|String|
|channel_id (PK)|String|
|date_created|Timestamp|
|date_updated|Timestamp|
|title|String|
|link|String|
|emoji|String|
|icon_url|String|
|type|String|
|rank|String|
|last_updated_by_user_id|String|
|last_updated_by_team_id|String|
|shortcut_id|String|
|entity_id|String|
|app_id|String|