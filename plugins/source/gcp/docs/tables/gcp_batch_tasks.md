# Table: gcp_batch_tasks

https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs.taskGroups.tasks/list

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_batch_task_groups](gcp_batch_task_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|status|JSON|