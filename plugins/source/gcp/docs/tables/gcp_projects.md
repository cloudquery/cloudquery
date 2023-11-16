# Table: gcp_projects

This table shows data for GCP Projects.

This table contains the list of all project_id's synced by cloudquery. It may contain projects missing from `gcp_resourcemanager_projects` (i.e. projects where the `resourcemanager` API is not enabled)

The primary key for this table is **project_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|