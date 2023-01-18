# Table: gcp_artifactregistry_repositories

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_locations](gcp_artifactregistry_locations.md).

The following tables depend on gcp_artifactregistry_repositories:
  - [gcp_artifactregistry_docker_images](gcp_artifactregistry_docker_images.md)
  - [gcp_artifactregistry_files](gcp_artifactregistry_files.md)
  - [gcp_artifactregistry_packages](gcp_artifactregistry_packages.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|format|String|
|description|String|
|labels|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|kms_key_name|String|