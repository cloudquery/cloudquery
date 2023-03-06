# Table: gcp_workflows_workflows

https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows#resource:-workflow

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
|description|String|
|state|String|
|revision_id|String|
|create_time|Timestamp|
|update_time|Timestamp|
|revision_create_time|Timestamp|
|labels|JSON|
|service_account|String|