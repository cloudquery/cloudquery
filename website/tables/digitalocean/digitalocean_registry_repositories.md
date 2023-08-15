# Table: digitalocean_registry_repositories

This table shows data for DigitalOcean Registry Repositories.

Deprecated. https://docs.digitalocean.com/reference/api/api-reference/#operation/registry_list_repositories

The primary key for this table is **name**.

## Relations

This table depends on [digitalocean_registries](digitalocean_registries).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|registry_name|`utf8`|
|latest_tag|`json`|
|tag_count|`int64`|