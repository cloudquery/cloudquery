# Table: gcp_projects

This table contains the list of all project_id's synced by cloudquery. It may contain projects missing from `gcp_resourcemanager_projects` (i.e. projects where the `resourcemanager` API is not enabled)

The primary key for this table is **project_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|