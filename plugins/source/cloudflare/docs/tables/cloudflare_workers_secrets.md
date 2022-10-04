# Table: cloudflare_workers_secrets


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`cloudflare_worker_meta_data`](cloudflare_worker_meta_data.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|worker_meta_data_id|String|
|name|String|
|type|String|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|