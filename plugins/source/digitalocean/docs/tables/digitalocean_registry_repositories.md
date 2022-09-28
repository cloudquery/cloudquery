# Table: digitalocean_registry_repositories


The primary key for this table is **name**.

## Relations
This table depends on [`digitalocean_registries`](digitalocean_registries.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|name (PK)|String|
|registry_name|String|
|latest_tag|JSON|
|tag_count|Int|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|