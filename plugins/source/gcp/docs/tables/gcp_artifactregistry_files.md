# Table: gcp_artifactregistry_files

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.files#File

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_artifactregistry_repositories](gcp_artifactregistry_repositories.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|size_bytes|Int|
|hashes|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|owner|String|