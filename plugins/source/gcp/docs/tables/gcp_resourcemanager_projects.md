# Table: gcp_resourcemanager_projects

This table shows data for GCP Resourcemanager Projects.

https://cloud.google.com/resource-manager/reference/rest/v3/projects#Project

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|parent|`utf8`|
|state|`utf8`|
|display_name|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|delete_time|`timestamp[us, tz=UTC]`|
|etag|`utf8`|
|labels|`json`|