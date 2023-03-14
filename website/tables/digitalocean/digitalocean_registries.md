# Table: digitalocean_registries

This table shows data for Digitalocean Registries.

The primary key for this table is **name**.

## Relations

The following tables depend on digitalocean_registries:
  - [digitalocean_registry_repositories](digitalocean_registry_repositories)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|storage_usage_bytes|Int|
|storage_usage_bytes_updated_at|Timestamp|
|created_at|Timestamp|
|region|String|