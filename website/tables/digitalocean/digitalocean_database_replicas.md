# Table: digitalocean_database_replicas

This table shows data for DigitalOcean Database Replicas.

https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_replicas

The primary key for this table is **_cq_id**.

## Relations

This table depends on [digitalocean_databases](digitalocean_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|id|`utf8`|
|name|`utf8`|
|connection|`json`|
|private_connection|`json`|
|region|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|private_network_uuid|`utf8`|
|tags|`list<item: utf8, nullable>`|