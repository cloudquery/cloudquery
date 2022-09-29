# Table: digitalocean_projects


The primary key for this table is **id**.

## Relations
The following tables depend on `digitalocean_projects`:
  - [`digitalocean_project_resources`](digitalocean_project_resources.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|id (PK)|String|
|owner_uuid|String|
|owner_id|Int|
|name|String|
|description|String|
|purpose|String|
|environment|String|
|is_default|Bool|
|created_at|String|
|updated_at|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|