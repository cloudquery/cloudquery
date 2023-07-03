# Table: gcp_clouddeploy_releases

This table shows data for GCP Clouddeploy Releases.

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases#Release

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_clouddeploy_delivery_pipelines](gcp_clouddeploy_delivery_pipelines).

The following tables depend on gcp_clouddeploy_releases:
  - [gcp_clouddeploy_rollouts](gcp_clouddeploy_rollouts)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|uid|`utf8`|
|description|`utf8`|
|annotations|`json`|
|labels|`json`|
|abandoned|`bool`|
|create_time|`timestamp[us, tz=UTC]`|
|render_start_time|`timestamp[us, tz=UTC]`|
|render_end_time|`timestamp[us, tz=UTC]`|
|skaffold_config_uri|`utf8`|
|skaffold_config_path|`utf8`|
|build_artifacts|`json`|
|delivery_pipeline_snapshot|`json`|
|target_snapshots|`json`|
|render_state|`utf8`|
|etag|`utf8`|
|skaffold_version|`utf8`|
|target_artifacts|`json`|
|target_renders|`json`|
|condition|`json`|