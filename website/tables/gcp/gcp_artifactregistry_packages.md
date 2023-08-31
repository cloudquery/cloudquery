# Table: gcp_artifactregistry_packages

This table shows data for GCP Artifact Registry Packages.

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages#Package

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_repositories](gcp_artifactregistry_repositories).

The following tables depend on gcp_artifactregistry_packages:
  - [gcp_artifactregistry_tags](gcp_artifactregistry_tags)
  - [gcp_artifactregistry_versions](gcp_artifactregistry_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|