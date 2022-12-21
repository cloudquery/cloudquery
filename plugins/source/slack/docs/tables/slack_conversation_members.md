# Table: slack_conversation_members

https://api.slack.com/methods/conversations.members

The composite primary key for this table is (**team_id**, **user_id**, **channel_id**).

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
|user_id (PK)|String|
|channel_id (PK)|String|