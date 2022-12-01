# Table: cloudflare_worker_meta_data



The primary key for this table is **id**.

## Relations

The following tables depend on cloudflare_worker_meta_data:
  - [cloudflare_worker_cron_triggers](cloudflare_worker_cron_triggers.md)
  - [cloudflare_workers_secrets](cloudflare_workers_secrets.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|id (PK)|String|
|etag|String|
|size|Int|
|created_on|Timestamp|
|modified_on|Timestamp|