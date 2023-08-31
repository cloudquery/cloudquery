# Table: digitalocean_registries

This table shows data for DigitalOcean Registries.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Container-Registry

The primary key for this table is **name**.

## Relations

The following tables depend on digitalocean_registries:
  - [digitalocean_registry_repositories](digitalocean_registry_repositories)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|storage_usage_bytes|`int64`|
|storage_usage_bytes_updated_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|region|`utf8`|