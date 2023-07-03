# Table: gcp_run_services

This table shows data for GCP Run Services.

https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#Service

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_run_locations](gcp_run_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|uid|`utf8`|
|generation|`int64`|
|labels|`json`|
|annotations|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|delete_time|`timestamp[us, tz=UTC]`|
|expire_time|`timestamp[us, tz=UTC]`|
|creator|`utf8`|
|last_modifier|`utf8`|
|client|`utf8`|
|client_version|`utf8`|
|ingress|`utf8`|
|launch_stage|`utf8`|
|binary_authorization|`json`|
|template|`json`|
|traffic|`json`|
|observed_generation|`int64`|
|terminal_condition|`json`|
|conditions|`json`|
|latest_ready_revision|`utf8`|
|latest_created_revision|`utf8`|
|traffic_statuses|`json`|
|uri|`utf8`|
|reconciling|`bool`|
|etag|`utf8`|