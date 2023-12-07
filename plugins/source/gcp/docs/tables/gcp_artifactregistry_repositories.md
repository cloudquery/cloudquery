# Table: gcp_artifactregistry_repositories

This table shows data for GCP Artifact Registry Repositories.

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
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|format_config|`json`|
|mode_config|`json`|
|name (PK)|`utf8`|
|format|`utf8`|
|description|`utf8`|
|labels|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|kms_key_name|`utf8`|
|mode|`utf8`|
|cleanup_policies|`json`|
|size_bytes|`int64`|
|satisfies_pzs|`bool`|
|cleanup_policy_dry_run|`bool`|