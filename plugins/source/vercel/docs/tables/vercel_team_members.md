# Table: vercel_team_members

The composite primary key for this table is (**team_id**, **uid**).

## Relations

This table depends on [vercel_teams](vercel_teams.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|team_id (PK)|String|
|avatar|String|
|confirmed|Bool|
|email|String|
|github|JSON|
|gitlab|JSON|
|bitbucket|JSON|
|role|String|
|uid (PK)|String|
|username|String|
|name|String|
|created_at|Timestamp|
|access_requested_at|Timestamp|