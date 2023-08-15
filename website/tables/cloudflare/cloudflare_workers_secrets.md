# Table: cloudflare_workers_secrets

This table shows data for Cloudflare Workers Secrets.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [cloudflare_worker_meta_data](cloudflare_worker_meta_data).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|worker_meta_data_id|`utf8`|
|name|`utf8`|
|secret_text|`utf8`|