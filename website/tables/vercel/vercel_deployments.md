# Table: vercel_deployments

This table shows data for Vercel Deployments.

The primary key for this table is **uid**.
It supports incremental syncs.
## Relations

The following tables depend on vercel_deployments:
  - [vercel_deployment_checks](vercel_deployment_checks)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|uid (PK)|`utf8`|
|name|`utf8`|
|url|`utf8`|
|source|`utf8`|
|state|`utf8`|
|type|`utf8`|
|inspector_url|`utf8`|
|is_rollback_candidate|`bool`|
|ready|`timestamp[us, tz=UTC]`|
|checks_state|`utf8`|
|checks_conclusion|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|building_at|`timestamp[us, tz=UTC]`|