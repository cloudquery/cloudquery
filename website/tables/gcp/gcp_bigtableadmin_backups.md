# Table: gcp_bigtableadmin_backups

This table shows data for GCP Bigtableadmin Backups.

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters.backups#Backup

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_clusters](gcp_bigtableadmin_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|source_table|`utf8`|
|size_bytes|`int64`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|expire_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|encryption_info|`json`|