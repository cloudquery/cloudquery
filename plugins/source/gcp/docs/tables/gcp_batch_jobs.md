# Table: gcp_batch_jobs

https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#Job

The composite primary key for this table is (**project_id**, **name**).

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
|priority|Int|
|task_groups|JSON|
|allocation_policy|JSON|
|labels|JSON|
|status|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|logs_policy|JSON|
|notifications|JSON|