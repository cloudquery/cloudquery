# Table: vercel_team_members

This table shows data for Vercel Team Members.

The composite primary key for this table is (**team_id**, **uid**).
It supports incremental syncs.
## Relations

This table depends on [vercel_teams](vercel_teams).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|team_id (PK)|`utf8`|
|uid (PK)|`utf8`|
|avatar|`utf8`|
|confirmed|`bool`|
|email|`utf8`|
|github|`json`|
|gitlab|`json`|
|bitbucket|`json`|
|role|`utf8`|
|username|`utf8`|
|name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|access_requested_at|`timestamp[us, tz=UTC]`|