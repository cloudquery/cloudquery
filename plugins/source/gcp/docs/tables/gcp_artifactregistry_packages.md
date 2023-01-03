# Table: gcp_artifactregistry_packages

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages#Package

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_repositories](gcp_artifactregistry_repositories.md).

The following tables depend on gcp_artifactregistry_packages:
  - [gcp_artifactregistry_tags](gcp_artifactregistry_tags.md)
  - [gcp_artifactregistry_versions](gcp_artifactregistry_versions.md)

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
|create_time|Timestamp|
|update_time|Timestamp|