# Table: digitalocean_database_firewall_rules

This table shows data for DigitalOcean Database Firewall Rules.

https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_firewall_rules

The primary key for this table is **_cq_id**.

## Relations

This table depends on [digitalocean_databases](digitalocean_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|uuid|`utf8`|
|cluster_uuid|`utf8`|
|type|`utf8`|
|value|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|