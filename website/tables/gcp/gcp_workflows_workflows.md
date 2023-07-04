# Table: gcp_workflows_workflows

This table shows data for GCP Workflows Workflows.

https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows#resource:-workflow

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|state|`utf8`|
|revision_id|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|revision_create_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|service_account|`utf8`|