# Table: gcp_functions_functions

This table shows data for GCP Functions Functions.

https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions#CloudFunction

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|status|`utf8`|
|entry_point|`utf8`|
|runtime|`utf8`|
|timeout|`int64`|
|available_memory_mb|`int64`|
|service_account_email|`utf8`|
|update_time|`timestamp[us, tz=UTC]`|
|version_id|`int64`|
|labels|`json`|
|environment_variables|`json`|
|build_environment_variables|`json`|
|network|`utf8`|
|max_instances|`int64`|
|min_instances|`int64`|
|vpc_connector|`utf8`|
|vpc_connector_egress_settings|`utf8`|
|ingress_settings|`utf8`|
|kms_key_name|`utf8`|
|build_worker_pool|`utf8`|
|build_id|`utf8`|
|build_name|`utf8`|
|secret_environment_variables|`json`|
|secret_volumes|`json`|
|source_token|`utf8`|
|docker_repository|`utf8`|
|docker_registry|`utf8`|