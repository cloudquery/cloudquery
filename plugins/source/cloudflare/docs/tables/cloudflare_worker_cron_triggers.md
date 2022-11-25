# Table: cloudflare_worker_cron_triggers



The primary key for this table is **_cq_id**.

## Relations
This table depends on [cloudflare_worker_meta_data](cloudflare_worker_meta_data.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|worker_meta_data_id|String|
|cron|String|
|created_on|Timestamp|
|modified_on|Timestamp|