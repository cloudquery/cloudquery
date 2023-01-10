# Table: gcp_clouddeploy_rollouts

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts#Rollout

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_clouddeploy_releases](gcp_clouddeploy_releases.md).

The following tables depend on gcp_clouddeploy_rollouts:
  - [gcp_clouddeploy_job_runs](gcp_clouddeploy_job_runs.md)

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
|create_time|Timestamp|
|approve_time|Timestamp|
|enqueue_time|Timestamp|
|deploy_start_time|Timestamp|
|deploy_end_time|Timestamp|
|target_id|String|
|approval_state|String|
|state|String|
|failure_reason|String|
|deploying_build|String|
|etag|String|
|deploy_failure_cause|String|
|phases|JSON|
|metadata|JSON|