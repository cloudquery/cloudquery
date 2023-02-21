# Table: vercel_deployments

The primary key for this table is **uid**.
It supports incremental syncs.
## Relations

The following tables depend on vercel_deployments:
  - [vercel_deployment_checks](vercel_deployment_checks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|uid (PK)|String|
|name|String|
|url|String|
|source|String|
|state|String|
|type|String|
|inspector_url|String|
|is_rollback_candidate|Bool|
|ready|Timestamp|
|checks_state|String|
|checks_conclusion|String|
|created_at|Timestamp|
|building_at|Timestamp|