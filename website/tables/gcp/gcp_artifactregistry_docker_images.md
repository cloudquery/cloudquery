# Table: gcp_artifactregistry_docker_images

This table shows data for GCP Artifact Registry Docker Images.

https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.dockerImages#DockerImage

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
|uri|`utf8`|
|tags|`list<item: utf8, nullable>`|
|image_size_bytes|`int64`|
|upload_time|`timestamp[us, tz=UTC]`|
|media_type|`utf8`|
|build_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|