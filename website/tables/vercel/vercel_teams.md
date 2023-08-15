# Table: vercel_teams

This table shows data for Vercel Teams.

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

The following tables depend on vercel_teams:
  - [vercel_team_members](vercel_team_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|slug|`utf8`|
|name|`utf8`|
|avatar|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|creator_id|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|profiles|`json`|
|staging_prefix|`utf8`|