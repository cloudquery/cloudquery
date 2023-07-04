# Table: gcp_bigtableadmin_app_profiles

This table shows data for GCP Bigtableadmin App Profiles.

https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.appProfiles#AppProfile

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_bigtableadmin_instances](gcp_bigtableadmin_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|etag|`utf8`|
|description|`utf8`|