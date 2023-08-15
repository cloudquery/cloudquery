# Table: cloudflare_worker_meta_data

This table shows data for Cloudflare Worker Meta Data.

The primary key for this table is **id**.

## Relations

The following tables depend on cloudflare_worker_meta_data:
  - [cloudflare_worker_cron_triggers](cloudflare_worker_cron_triggers)
  - [cloudflare_workers_secrets](cloudflare_workers_secrets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id (PK)|`utf8`|
|etag|`utf8`|
|size|`int64`|
|created_on|`timestamp[us, tz=UTC]`|
|modified_on|`timestamp[us, tz=UTC]`|