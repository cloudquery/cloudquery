# Table: gcp_artifactregistry_files

This table shows data for GCP Artifact Registry Files.

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.files#File

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_repositories](gcp_artifactregistry_repositories).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|size_bytes|`int64`|
|hashes|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|owner|`utf8`|
|fetch_time|`timestamp[us, tz=UTC]`|