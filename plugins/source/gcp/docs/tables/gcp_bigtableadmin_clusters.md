# Table: gcp_bigtableadmin_clusters

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters#Cluster

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_instances](gcp_bigtableadmin_instances.md).

The following tables depend on gcp_bigtableadmin_clusters:
  - [gcp_bigtableadmin_backups](gcp_bigtableadmin_backups.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|zone|String|
|serve_nodes|Int|
|state|String|
|storage_type|Int|
|kms_key_name|String|
|autoscaling_config|JSON|