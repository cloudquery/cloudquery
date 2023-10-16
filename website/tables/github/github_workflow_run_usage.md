# Table: github_workflow_run_usage

This table shows data for Github Workflow Run Usage.

The composite primary key for this table is (**org**, **repository_id**, **run_id**).

## Relations

This table depends on [github_workflow_runs](github_workflow_runs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|run_id (PK)|`int64`|
|billable|`json`|
|run_duration_ms|`int64`|