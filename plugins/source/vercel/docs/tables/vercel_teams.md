# Table: vercel_teams

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

The following tables depend on vercel_teams:
  - [vercel_team_members](vercel_team_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|slug|String|
|name|String|
|avatar|String|
|created_at|Timestamp|
|creator_id|String|
|updated_at|Timestamp|
|profiles|JSON|
|staging_prefix|String|