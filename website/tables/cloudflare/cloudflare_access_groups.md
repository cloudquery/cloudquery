# Table: cloudflare_access_groups

This table shows data for Cloudflare Access Groups.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|zone_id|utf8|
|id (PK)|utf8|
|created_at|timestamp[us, tz=UTC]|
|updated_at|timestamp[us, tz=UTC]|
|name|utf8|
|include|json|
|exclude|json|
|require|json|