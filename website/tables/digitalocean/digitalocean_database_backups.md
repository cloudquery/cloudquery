# Table: digitalocean_database_backups

This table shows data for DigitalOcean Database Backups.

https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_backups

The primary key for this table is **_cq_id**.

## Relations

This table depends on [digitalocean_databases](digitalocean_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|created_at|`timestamp[us, tz=UTC]`|
|size_gigabytes|`float64`|