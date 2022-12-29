# Table: gcp_artifactregistry_locations

https://cloud.google.com/artifact-registry/docs/reference/rest/Shared.Types/ListLocationsResponse#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_artifactregistry_locations:
  - [gcp_artifactregistry_repositories](gcp_artifactregistry_repositories.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|display_name|String|
|labels|JSON|
|location_id|String|
|metadata|IntArray|
|name (PK)|String|