# Table: digitalocean_databases

This table shows data for DigitalOcean Databases.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_databases:
  - [digitalocean_database_backups](digitalocean_database_backups)
  - [digitalocean_database_firewall_rules](digitalocean_database_firewall_rules)
  - [digitalocean_database_replicas](digitalocean_database_replicas)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|engine|`utf8`|
|version|`utf8`|
|connection|`json`|
|private_connection|`json`|
|users|`json`|
|num_nodes|`int64`|
|size|`utf8`|
|db_names|`list<item: utf8, nullable>`|
|region|`utf8`|
|status|`utf8`|
|maintenance_window|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|private_network_uuid|`utf8`|
|tags|`list<item: utf8, nullable>`|
|project_id|`utf8`|