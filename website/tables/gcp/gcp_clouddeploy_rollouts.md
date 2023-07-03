# Table: gcp_clouddeploy_rollouts

This table shows data for GCP Clouddeploy Rollouts.

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts#Rollout

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_clouddeploy_releases](gcp_clouddeploy_releases).

The following tables depend on gcp_clouddeploy_rollouts:
  - [gcp_clouddeploy_job_runs](gcp_clouddeploy_job_runs)

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
|create_time|`timestamp[us, tz=UTC]`|
|approve_time|`timestamp[us, tz=UTC]`|
|enqueue_time|`timestamp[us, tz=UTC]`|
|deploy_start_time|`timestamp[us, tz=UTC]`|
|deploy_end_time|`timestamp[us, tz=UTC]`|
|target_id|`utf8`|
|approval_state|`utf8`|
|state|`utf8`|
|failure_reason|`utf8`|
|deploying_build|`utf8`|
|etag|`utf8`|
|deploy_failure_cause|`utf8`|
|phases|`json`|
|metadata|`json`|
|controller_rollout|`utf8`|