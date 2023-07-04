# Table: gcp_artifactregistry_locations

This table shows data for GCP Artifact Registry Locations.

https://cloud.google.com/artifact-registry/docs/reference/rest/Shared.Types/ListLocationsResponse#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_artifactregistry_locations:
  - [gcp_artifactregistry_repositories](gcp_artifactregistry_repositories)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|location_id|`utf8`|
|metadata|`binary`|
|name (PK)|`utf8`|