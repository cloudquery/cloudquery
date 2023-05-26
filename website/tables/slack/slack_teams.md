# Table: slack_teams

This table shows data for Slack Teams.

https://slack.com/api/team.info

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|name|utf8|
|domain|utf8|
|email_domain|utf8|
|icon|json|