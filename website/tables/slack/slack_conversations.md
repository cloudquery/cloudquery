# Table: slack_conversations

This table shows data for Slack Conversations.

https://api.slack.com/methods/conversations.list

The composite primary key for this table is (**team_id**, **id**).

## Relations

The following tables depend on slack_conversations:
  - [slack_conversation_bookmarks](slack_conversation_bookmarks)
  - [slack_conversation_histories](slack_conversation_histories)
  - [slack_conversation_members](slack_conversation_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|team_id (PK)|utf8|
|id (PK)|utf8|
|created|timestamp[us, tz=UTC]|
|is_open|bool|
|last_read|utf8|
|unread_count|int64|
|unread_count_display|int64|
|is_group|bool|
|is_shared|bool|
|is_im|bool|
|is_ext_shared|bool|
|is_org_shared|bool|
|is_pending_ext_shared|bool|
|is_private|bool|
|is_mpim|bool|
|unlinked|int64|
|name_normalized|utf8|
|num_members|int64|
|priority|float64|
|user|utf8|
|name|utf8|
|creator|utf8|
|is_archived|bool|
|topic|json|
|purpose|json|
|is_channel|bool|
|is_general|bool|
|is_member|bool|
|locale|utf8|