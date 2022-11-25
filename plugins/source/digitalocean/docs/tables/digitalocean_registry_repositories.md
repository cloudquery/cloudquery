# Table: digitalocean_registry_repositories



The primary key for this table is **name**.

## Relations
This table depends on [digitalocean_registries](digitalocean_registries.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|registry_name|String|
|latest_tag|JSON|
|tag_count|Int|