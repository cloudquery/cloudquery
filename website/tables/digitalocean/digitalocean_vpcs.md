# Table: digitalocean_vpcs

This table shows data for DigitalOcean VPCs.

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_vpcs:
  - [digitalocean_vpc_members](digitalocean_vpc_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|urn|String|
|name|String|
|description|String|
|ip_range|String|
|region|String|
|created_at|Timestamp|
|default|Bool|