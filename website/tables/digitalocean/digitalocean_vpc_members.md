# Table: digitalocean_vpc_members

This table shows data for DigitalOcean VPC Members.

https://docs.digitalocean.com/reference/api/api-reference/#operation/vpcs_list_members

The primary key for this table is **urn**.

## Relations

This table depends on [digitalocean_vpcs](digitalocean_vpcs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|urn (PK)|`utf8`|
|name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|