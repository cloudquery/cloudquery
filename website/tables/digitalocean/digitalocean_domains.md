# Table: digitalocean_domains

This table shows data for DigitalOcean Domains.

https://docs.digitalocean.com/reference/api/api-reference/#operation/domains_list

The primary key for this table is **name**.

## Relations

The following tables depend on digitalocean_domains:
  - [digitalocean_domain_records](digitalocean_domain_records)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|ttl|`int64`|
|zone_file|`utf8`|