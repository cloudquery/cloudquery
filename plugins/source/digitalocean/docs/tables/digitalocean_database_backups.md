# Table: digitalocean_database_backups



The primary key for this table is **_cq_id**.

## Relations
This table depends on [digitalocean_databases](digitalocean_databases.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|created_at|Timestamp|
|size_gigabytes|Float|