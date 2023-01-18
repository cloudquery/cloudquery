# Table: gcp_clouddeploy_releases

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases#Release

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_clouddeploy_delivery_pipelines](gcp_clouddeploy_delivery_pipelines.md).

The following tables depend on gcp_clouddeploy_releases:
  - [gcp_clouddeploy_rollouts](gcp_clouddeploy_rollouts.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|uid|String|
|description|String|
|annotations|JSON|
|labels|JSON|
|abandoned|Bool|
|create_time|Timestamp|
|render_start_time|Timestamp|
|render_end_time|Timestamp|
|skaffold_config_uri|String|
|skaffold_config_path|String|
|build_artifacts|JSON|
|delivery_pipeline_snapshot|JSON|
|target_snapshots|JSON|
|render_state|String|
|etag|String|
|skaffold_version|String|
|target_artifacts|JSON|
|target_renders|JSON|