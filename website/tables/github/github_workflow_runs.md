# Table: github_workflow_runs

This table shows data for Github Workflow Runs.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

The following tables depend on github_workflow_runs:
  - [github_workflow_jobs](github_workflow_jobs)
  - [github_workflow_run_usage](github_workflow_run_usage)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|id (PK)|`int64`|
|name|`utf8`|
|node_id|`utf8`|
|head_branch|`utf8`|
|head_sha|`utf8`|
|run_number|`int64`|
|run_attempt|`int64`|
|event|`utf8`|
|status|`utf8`|
|conclusion|`utf8`|
|workflow_id|`int64`|
|check_suite_id|`int64`|
|check_suite_node_id|`utf8`|
|url|`utf8`|
|html_url|`utf8`|
|pull_requests|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|run_started_at|`timestamp[us, tz=UTC]`|
|jobs_url|`utf8`|
|logs_url|`utf8`|
|check_suite_url|`utf8`|
|artifacts_url|`utf8`|
|cancel_url|`utf8`|
|rerun_url|`utf8`|
|previous_attempt_url|`utf8`|
|head_commit|`json`|
|workflow_url|`utf8`|
|repository|`json`|
|head_repository|`json`|
|actor|`json`|