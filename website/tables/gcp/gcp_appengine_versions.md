# Table: gcp_appengine_versions

This table shows data for GCP App Engine Versions.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_appengine_services](gcp_appengine_services).

The following tables depend on gcp_appengine_versions:
  - [gcp_appengine_instances](gcp_appengine_instances)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|inbound_services|`list<item: int64, nullable>`|
|instance_class|`utf8`|
|network|`json`|
|zones|`list<item: utf8, nullable>`|
|resources|`json`|
|runtime|`utf8`|
|runtime_channel|`utf8`|
|threadsafe|`bool`|
|vm|`bool`|
|app_engine_apis|`bool`|
|beta_settings|`json`|
|env|`utf8`|
|serving_status|`utf8`|
|created_by|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|disk_usage_bytes|`int64`|
|runtime_api_version|`utf8`|
|runtime_main_executable_path|`utf8`|
|service_account|`utf8`|
|handlers|`json`|
|error_handlers|`json`|
|libraries|`json`|
|api_config|`json`|
|env_variables|`json`|
|build_env_variables|`json`|
|default_expiration|`int64`|
|health_check|`json`|
|readiness_check|`json`|
|liveness_check|`json`|
|nobuild_files_regex|`utf8`|
|deployment|`json`|
|version_url|`utf8`|
|endpoints_api_service|`json`|
|entrypoint|`json`|
|vpc_access_connector|`json`|