# Table: digitalocean_spaces



The primary key for this table is **_cq_id**.

## Relations

The following tables depend on digitalocean_spaces:
  - [digitalocean_space_cors](digitalocean_space_cors.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|acls|JSON|
|bucket|JSON|
|location|String|
|public|Bool|