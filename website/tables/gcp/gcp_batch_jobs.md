# Table: gcp_batch_jobs

This table shows data for GCP Batch Jobs.

https://cloud.google.com/batch/docs/reference/rest/v1/projects.locations.jobs#Job

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|uid|`utf8`|
|priority|`int64`|
|task_groups|`json`|
|allocation_policy|`json`|
|labels|`json`|
|status|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|logs_policy|`json`|
|notifications|`json`|