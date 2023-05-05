# Table: digitalocean_vpc_members

This table shows data for DigitalOcean VPC Members.

https://docs.digitalocean.com/reference/api/api-reference/#operation/vpcs_list_members

The primary key for this table is **urn**.

## Relations

This table depends on [digitalocean_vpcs](digitalocean_vpcs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|urn (PK)|String|
|name|String|
|created_at|Timestamp|