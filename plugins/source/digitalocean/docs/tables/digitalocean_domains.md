# Table: digitalocean_domains


The primary key for this table is **name**.

## Relations
The following tables depend on `digitalocean_domains`:
  - [`digitalocean_domain_records`](digitalocean_domain_records.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|name (PK)|String|
|ttl|Int|
|zone_file|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|