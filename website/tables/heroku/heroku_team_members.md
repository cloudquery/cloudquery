# Table: heroku_team_members

This table shows data for Heroku Team Members.

https://devcenter.heroku.com/articles/platform-api-reference#team-member

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|email|`utf8`|
|federated|`bool`|
|identity_provider|`json`|
|role|`utf8`|
|two_factor_authentication|`bool`|
|updated_at|`timestamp[us, tz=UTC]`|
|user|`json`|