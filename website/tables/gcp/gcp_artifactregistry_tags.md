# Table: gcp_artifactregistry_tags

This table shows data for GCP Artifact Registry Tags.

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.tags#Tag

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
|version|`utf8`|