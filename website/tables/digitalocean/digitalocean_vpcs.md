# Table: digitalocean_vpcs

This table shows data for DigitalOcean VPCs.

https://docs.digitalocean.com/reference/api/api-reference/#tag/VPCs

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_vpcs:
  - [digitalocean_vpc_members](digitalocean_vpc_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|urn|`utf8`|
|name|`utf8`|
|description|`utf8`|
|ip_range|`utf8`|
|region|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|default|`bool`|