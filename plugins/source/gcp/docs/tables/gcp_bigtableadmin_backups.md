# Table: gcp_bigtableadmin_backups

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters.backups#Backup

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_clusters](gcp_bigtableadmin_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|source_table|String|
|size_bytes|Int|
|start_time|Timestamp|
|end_time|Timestamp|
|expire_time|Timestamp|
|state|String|
|encryption_info|JSON|