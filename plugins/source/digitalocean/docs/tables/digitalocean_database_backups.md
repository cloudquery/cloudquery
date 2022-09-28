# Table: digitalocean_database_backups


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`digitalocean_databases`](digitalocean_databases.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|created_at|Timestamp|
|size_gigabytes|Float|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|