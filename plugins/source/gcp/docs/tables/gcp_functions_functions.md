# Table: gcp_functions_functions

https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions#CloudFunction

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
|status|String|
|entry_point|String|
|runtime|String|
|timeout|Int|
|available_memory_mb|Int|
|service_account_email|String|
|update_time|Timestamp|
|version_id|Int|
|labels|JSON|
|environment_variables|JSON|
|build_environment_variables|JSON|
|network|String|
|max_instances|Int|
|min_instances|Int|
|vpc_connector|String|
|vpc_connector_egress_settings|String|
|ingress_settings|String|
|kms_key_name|String|
|build_worker_pool|String|
|build_id|String|
|build_name|String|
|secret_environment_variables|JSON|
|secret_volumes|JSON|
|source_token|String|
|docker_repository|String|
|docker_registry|String|