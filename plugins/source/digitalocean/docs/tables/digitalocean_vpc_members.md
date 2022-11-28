# Table: digitalocean_vpc_members



The primary key for this table is **urn**.

## Relations
This table depends on [digitalocean_vpcs](digitalocean_vpcs.md).


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