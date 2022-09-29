# Table: digitalocean_vpcs


The primary key for this table is **id**.

## Relations
The following tables depend on `digitalocean_vpcs`:
  - [`digitalocean_vpc_members`](digitalocean_vpc_members.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|urn|String|
|name|String|
|description|String|
|ip_range|String|
|region|String|
|created_at|Timestamp|
|default|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|