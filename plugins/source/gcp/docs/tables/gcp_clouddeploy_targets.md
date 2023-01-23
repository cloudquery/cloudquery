# Table: gcp_clouddeploy_targets

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.targets#Target

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
|target_id|String|
|uid|String|
|description|String|
|annotations|JSON|
|labels|JSON|
|require_approval|Bool|
|create_time|Timestamp|
|update_time|Timestamp|
|etag|String|
|execution_configs|JSON|