# Table: gcp_bigtableadmin_app_profiles

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.appProfiles#AppProfile

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_instances](gcp_bigtableadmin_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|etag|String|
|description|String|