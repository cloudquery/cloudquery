# Table: vercel_deployment_checks

This table shows data for Vercel Deployment Checks.

The composite primary key for this table is (**deployment_id**, **id**).
It supports incremental syncs.
## Relations

This table depends on [vercel_deployments](vercel_deployments).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|deployment_id (PK)|String|
|id (PK)|String|
|created_at|Timestamp|
|completed_at|Timestamp|
|conclusion|String|
|details_url|String|
|integration_id|String|
|name|String|
|path|String|
|rererequestable|Bool|
|started_at|Timestamp|
|updated_at|Timestamp|
|status|String|