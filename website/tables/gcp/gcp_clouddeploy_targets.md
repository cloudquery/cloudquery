# Table: gcp_clouddeploy_targets

This table shows data for GCP Clouddeploy Targets.

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.targets#Target

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|target_id|`utf8`|
|uid|`utf8`|
|description|`utf8`|
|annotations|`json`|
|labels|`json`|
|require_approval|`bool`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|etag|`utf8`|
|execution_configs|`json`|