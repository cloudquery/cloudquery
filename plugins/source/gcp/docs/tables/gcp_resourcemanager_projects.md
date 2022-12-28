# Table: gcp_resourcemanager_projects

https://cloud.google.com/resource-manager/reference/rest/v3/projects#Project

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
|parent|String|
|state|String|
|display_name|String|
|create_time|Timestamp|
|update_time|Timestamp|
|delete_time|Timestamp|
|etag|String|
|labels|JSON|