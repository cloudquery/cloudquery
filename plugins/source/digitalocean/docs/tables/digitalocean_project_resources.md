# Table: digitalocean_project_resources



The primary key for this table is **urn**.

## Relations
This table depends on [digitalocean_projects](digitalocean_projects.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|urn (PK)|String|
|assigned_at|String|
|links|JSON|
|status|String|