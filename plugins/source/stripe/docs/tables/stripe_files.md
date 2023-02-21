# Table: stripe_files

https://stripe.com/docs/api/files

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
|expires_at|Int|
|filename|String|
|links|JSON|
|object|String|
|purpose|String|
|size|Int|
|title|String|
|type|String|
|url|String|