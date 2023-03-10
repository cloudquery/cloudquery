# Table: stripe_file_links

https://stripe.com/docs/api/file_links

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|expired|Bool|
|expires_at|Int|
|file|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|url|String|