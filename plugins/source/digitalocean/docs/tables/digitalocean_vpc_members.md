# Table: digitalocean_vpc_members


The primary key for this table is **urn**.

## Relations
This table depends on [`digitalocean_vpcs`](digitalocean_vpcs.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|urn (PK)|String|
|name|String|
|created_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|