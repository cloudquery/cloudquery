# Table: vercel_domain_records

The composite primary key for this table is (**domain_name**, **id**).
It supports incremental syncs.
## Relations

This table depends on [vercel_domains](vercel_domains.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|domain_name (PK)|String|
|id (PK)|String|
|slug|String|
|name|String|
|type|String|
|value|String|
|mx_priority|Int|
|priority|Int|
|creator|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|ttl|Int|