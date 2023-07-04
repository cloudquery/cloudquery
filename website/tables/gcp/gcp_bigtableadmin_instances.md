# Table: gcp_bigtableadmin_instances

This table shows data for GCP Bigtableadmin Instances.

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances#Instance

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_bigtableadmin_instances:
  - [gcp_bigtableadmin_app_profiles](gcp_bigtableadmin_app_profiles)
  - [gcp_bigtableadmin_clusters](gcp_bigtableadmin_clusters)
  - [gcp_bigtableadmin_tables](gcp_bigtableadmin_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|instance_state|`int64`|
|instance_type|`int64`|
|labels|`json`|