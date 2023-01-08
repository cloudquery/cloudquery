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
|conversation|JSON|
|name|String|
|creator|String|
|is_archived|Bool|
|members|StringArray|
|topic|JSON|
|purpose|JSON|
|is_channel|Bool|
|is_general|Bool|
|is_member|Bool|
|locale|String|