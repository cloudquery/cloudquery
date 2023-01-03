# Table: gcp_batch_task_groups

https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#TaskGroup

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_batch_task_groups:
  - [gcp_batch_tasks](gcp_batch_tasks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|task_spec|JSON|
|task_count|Int|
|parallelism|Int|
|task_environments|JSON|
|task_count_per_node|Int|
|require_hosts_file|Bool|
|permissive_ssh|Bool|