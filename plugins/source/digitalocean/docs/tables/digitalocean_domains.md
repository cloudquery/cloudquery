# Table: digitalocean_domains



The primary key for this table is **name**.

## Relations
The following tables depend on `digitalocean_domains`:
  - [`digitalocean_domain_records`](digitalocean_domain_records.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|name (PK)|String|
|ttl|Int|
|zone_file|String|