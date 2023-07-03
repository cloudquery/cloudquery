# Table: gcp_artifactregistry_versions

This table shows data for GCP Artifact Registry Versions.

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.versions#Version

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_packages](gcp_artifactregistry_packages).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|related_tags|`json`|
|metadata|`json`|