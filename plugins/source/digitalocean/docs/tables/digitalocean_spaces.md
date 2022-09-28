# Table: digitalocean_spaces


The primary key for this table is **_cq_id**.

## Relations
The following tables depend on `digitalocean_spaces`:
  - [`digitalocean_space_cors`](digitalocean_space_cors.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|acls|JSON|
|bucket|JSON|
|location|String|
|public|Bool|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|