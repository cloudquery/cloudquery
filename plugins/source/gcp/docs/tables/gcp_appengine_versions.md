# Table: gcp_appengine_versions

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_appengine_services](gcp_appengine_services.md).

The following tables depend on gcp_appengine_versions:
  - [gcp_appengine_instances](gcp_appengine_instances.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|inbound_services|IntArray|
|instance_class|String|
|network|JSON|
|zones|StringArray|
|resources|JSON|
|runtime|String|
|runtime_channel|String|
|threadsafe|Bool|
|vm|Bool|
|app_engine_apis|Bool|
|beta_settings|JSON|
|env|String|
|serving_status|String|
|created_by|String|
|create_time|Timestamp|
|disk_usage_bytes|Int|
|runtime_api_version|String|
|runtime_main_executable_path|String|
|service_account|String|
|handlers|JSON|
|error_handlers|JSON|
|libraries|JSON|
|api_config|JSON|
|env_variables|JSON|
|build_env_variables|JSON|
|default_expiration|Int|
|health_check|JSON|
|readiness_check|JSON|
|liveness_check|JSON|
|nobuild_files_regex|String|
|deployment|JSON|
|version_url|String|
|endpoints_api_service|JSON|
|entrypoint|JSON|
|vpc_access_connector|JSON|