# Table: gcp_artifactregistry_docker_images

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.dockerImages#DockerImage

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
|uri|String|
|tags|StringArray|
|image_size_bytes|Int|
|upload_time|Timestamp|
|media_type|String|
|build_time|Timestamp|