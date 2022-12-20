# Table: slack_conversation_histories

https://api.slack.com/methods/conversations.history

The composite primary key for this table is (**channel_id**, **team_id**, **ts**).

## Relations

This table depends on [slack_conversations](slack_conversations.md).

The following tables depend on slack_conversation_histories:
  - [slack_conversation_replies](slack_conversation_replies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|channel_id (PK)|String|
|team_id (PK)|String|
|client_msg_id|String|
|type|String|
|user|String|
|text|String|
|ts (PK)|String|
|thread_ts|String|
|is_starred|Bool|
|pinned_to|StringArray|
|attachments|JSON|
|edited|JSON|
|last_read|String|
|subscribed|Bool|
|unread_count|Int|
|subtype|String|
|hidden|Bool|
|deleted_ts|String|
|event_ts|String|
|bot_id|String|
|username|String|
|icons|JSON|
|bot_profile|JSON|
|inviter|String|
|topic|String|
|purpose|String|
|name|String|
|old_name|String|
|members|StringArray|
|reply_count|Int|
|parent_user_id|String|
|latest_reply|String|
|files|JSON|
|upload|Bool|
|comment|JSON|
|item_type|String|
|reply_to|Int|
|reactions|JSON|
|response_type|String|
|replace_original|Bool|
|delete_original|Bool|
|metadata|JSON|
|permalink|String|