# Table: slack_conversations

https://api.slack.com/methods/conversations.list

The composite primary key for this table is (**team_id**, **id**).

## Relations

The following tables depend on slack_conversations:
  - [slack_conversation_bookmarks](slack_conversation_bookmarks.md)
  - [slack_conversation_histories](slack_conversation_histories.md)
  - [slack_conversation_members](slack_conversation_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|team_id (PK)|String|
|id (PK)|String|
|created|Timestamp|
|is_open|Bool|
|last_read|String|
|unread_count|Int|
|unread_count_display|Int|
|is_group|Bool|
|is_shared|Bool|
|is_im|Bool|
|is_ext_shared|Bool|
|is_org_shared|Bool|
|is_pending_ext_shared|Bool|
|is_private|Bool|
|is_mpim|Bool|
|unlinked|Int|
|name_normalized|String|
|num_members|Int|
|priority|Float|
|user|String|
|name|String|
|creator|String|
|is_archived|Bool|
|topic|JSON|
|purpose|JSON|
|is_channel|Bool|
|is_general|Bool|
|is_member|Bool|
|locale|String|