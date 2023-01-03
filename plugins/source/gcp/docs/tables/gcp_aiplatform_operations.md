# Table: gcp_aiplatform_operations

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.operations#Operation

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
|metadata|JSON|
|done|Bool|