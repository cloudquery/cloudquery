# Table: gcp_bigtableadmin_clusters

This table shows data for GCP Bigtableadmin Clusters.

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters#Cluster

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_instances](gcp_bigtableadmin_instances).

The following tables depend on gcp_bigtableadmin_clusters:
  - [gcp_bigtableadmin_backups](gcp_bigtableadmin_backups)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|zone|`utf8`|
|serve_nodes|`int64`|
|state|`utf8`|
|storage_type|`int64`|
|kms_key_name|`utf8`|
|autoscaling_config|`json`|