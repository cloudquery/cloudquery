# Table: gcp_artifactregistry_versions

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.versions#Version

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_packages](gcp_artifactregistry_packages.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|description|String|
|create_time|Timestamp|
|update_time|Timestamp|
|related_tags|JSON|
|metadata|JSON|