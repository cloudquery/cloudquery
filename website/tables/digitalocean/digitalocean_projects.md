# Table: digitalocean_projects

This table shows data for DigitalOcean Projects.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Projects

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_projects:
  - [digitalocean_project_resources](digitalocean_project_resources)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|owner_uuid|`utf8`|
|owner_id|`int64`|
|name|`utf8`|
|description|`utf8`|
|purpose|`utf8`|
|environment|`utf8`|
|is_default|`bool`|
|created_at|`utf8`|
|updated_at|`utf8`|