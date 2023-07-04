# Table: gcp_batch_tasks

This table shows data for GCP Batch Tasks.

https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs.taskGroups.tasks/list

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_batch_task_groups](gcp_batch_task_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|status|`json`|