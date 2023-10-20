# Table: github_workflow_jobs

This table shows data for Github Workflow Jobs.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

This table depends on [github_workflow_runs](github_workflow_runs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|run_id|`int64`|
|id (PK)|`int64`|
|run_url|`utf8`|
|node_id|`utf8`|
|head_sha|`utf8`|
|url|`utf8`|
|html_url|`utf8`|
|status|`utf8`|
|conclusion|`utf8`|
|started_at|`timestamp[us, tz=UTC]`|
|completed_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|steps|`json`|
|check_run_url|`utf8`|
|labels|`list<item: utf8, nullable>`|
|runner_id|`int64`|
|runner_name|`utf8`|
|runner_group_id|`int64`|
|runner_group_name|`utf8`|
|run_attempt|`int64`|