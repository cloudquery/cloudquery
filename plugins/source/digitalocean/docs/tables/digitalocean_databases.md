# Table: digitalocean_databases



The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_databases:
  - [digitalocean_database_firewall_rules](digitalocean_database_firewall_rules.md)
  - [digitalocean_database_replicas](digitalocean_database_replicas.md)
  - [digitalocean_database_backups](digitalocean_database_backups.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|engine|String|
|version|String|
|connection|JSON|
|private_connection|JSON|
|users|JSON|
|num_nodes|Int|
|size|String|
|db_names|StringArray|
|region|String|
|status|String|
|maintenance_window|JSON|
|created_at|Timestamp|
|private_network_uuid|String|
|tags|StringArray|
|project_id|String|