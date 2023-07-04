# Table: gcp_aiplatform_operations

This table shows data for GCP AI Platform Operations.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.operations#Operation

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|metadata|`json`|
|done|`bool`|