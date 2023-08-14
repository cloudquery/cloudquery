# Table: cloudflare_worker_cron_triggers

This table shows data for Cloudflare Worker Cron Triggers.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [cloudflare_worker_meta_data](cloudflare_worker_meta_data).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|worker_meta_data_id|`utf8`|
|cron|`utf8`|
|created_on|`timestamp[us, tz=UTC]`|
|modified_on|`timestamp[us, tz=UTC]`|