# Table: gcp_bigtableadmin_tables

This table shows data for GCP Bigtableadmin Tables.

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.tables#Table

The composite primary key for this table is (**project_id**, **instance_name**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_instances](gcp_bigtableadmin_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|instance_name (PK)|`utf8`|
|families|`list<item: utf8, nullable>`|
|family_infos|`json`|
|deletion_protection|`int64`|
|name (PK)|`utf8`|