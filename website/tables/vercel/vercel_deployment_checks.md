# Table: vercel_deployment_checks

This table shows data for Vercel Deployment Checks.

The composite primary key for this table is (**deployment_id**, **id**).
It supports incremental syncs.
## Relations

This table depends on [vercel_deployments](vercel_deployments).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|deployment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|completed_at|`timestamp[us, tz=UTC]`|
|conclusion|`utf8`|
|details_url|`utf8`|
|integration_id|`utf8`|
|name|`utf8`|
|path|`utf8`|
|rererequestable|`bool`|
|started_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|status|`utf8`|