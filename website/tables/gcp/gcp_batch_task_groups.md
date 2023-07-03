# Table: gcp_batch_task_groups

This table shows data for GCP Batch Task Groups.

https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#TaskGroup

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_batch_task_groups:
  - [gcp_batch_tasks](gcp_batch_tasks)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|task_spec|`json`|
|task_count|`int64`|
|parallelism|`int64`|
|task_environments|`json`|
|task_count_per_node|`int64`|
|require_hosts_file|`bool`|
|permissive_ssh|`bool`|