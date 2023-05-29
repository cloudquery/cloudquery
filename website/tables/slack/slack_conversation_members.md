# Table: slack_conversation_members

This table shows data for Slack Conversation Members.

https://api.slack.com/methods/conversations.members

The composite primary key for this table is (**team_id**, **user_id**, **channel_id**).

## Relations

This table depends on [slack_conversations](slack_conversations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|team_id (PK)|`utf8`|
|user_id (PK)|`utf8`|
|channel_id (PK)|`utf8`|