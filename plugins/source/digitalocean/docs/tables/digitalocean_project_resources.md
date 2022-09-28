# Table: digitalocean_project_resources


The primary key for this table is **urn**.

## Relations
This table depends on [`digitalocean_projects`](digitalocean_projects.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|urn (PK)|String|
|assigned_at|String|
|links|JSON|
|status|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|