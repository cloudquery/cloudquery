# Table: gcp_websecurityscanner_scan_configs

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs#resource:-scanconfig

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_websecurityscanner_scan_configs:
  - [gcp_websecurityscanner_scan_config_scan_runs](gcp_websecurityscanner_scan_config_scan_runs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|max_qps|Int|
|starting_urls|StringArray|
|authentication|JSON|
|user_agent|String|
|blacklist_patterns|StringArray|
|schedule|JSON|
|export_to_security_command_center|String|
|risk_level|String|
|managed_scan|Bool|
|static_ip_scan|Bool|
|ignore_http_status_errors|Bool|