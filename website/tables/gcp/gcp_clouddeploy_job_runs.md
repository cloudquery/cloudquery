# Table: gcp_clouddeploy_job_runs

This table shows data for GCP Clouddeploy Job Runs.

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines.releases.rollouts.jobRuns#JobRun

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_clouddeploy_rollouts](gcp_clouddeploy_rollouts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|uid|`utf8`|
|phase_id|`utf8`|
|job_id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|etag|`utf8`|