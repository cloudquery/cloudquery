# Table: digitalocean_project_resources

This table shows data for DigitalOcean Project Resources.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Project-Resources

The primary key for this table is **urn**.

## Relations

This table depends on [digitalocean_projects](digitalocean_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|urn (PK)|`utf8`|
|assigned_at|`utf8`|
|links|`json`|
|status|`utf8`|