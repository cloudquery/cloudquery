# Table: gcp_bigtableadmin_instances

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_bigtableadmin_instances:
  - [gcp_bigtableadmin_app_profiles](gcp_bigtableadmin_app_profiles.md)
  - [gcp_bigtableadmin_clusters](gcp_bigtableadmin_clusters.md)
  - [gcp_bigtableadmin_tables](gcp_bigtableadmin_tables.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|instance_state|Int|
|instance_type|Int|
|labels|JSON|