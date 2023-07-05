# Table: gcp_websecurityscanner_scan_configs

This table shows data for GCP Web Security Scanner Scan Configs.

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs#resource:-scanconfig

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_websecurityscanner_scan_configs:
  - [gcp_websecurityscanner_scan_config_scan_runs](gcp_websecurityscanner_scan_config_scan_runs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|max_qps|`int64`|
|starting_urls|`list<item: utf8, nullable>`|
|authentication|`json`|
|user_agent|`utf8`|
|blacklist_patterns|`list<item: utf8, nullable>`|
|schedule|`json`|
|export_to_security_command_center|`utf8`|
|risk_level|`utf8`|
|managed_scan|`bool`|
|static_ip_scan|`bool`|
|ignore_http_status_errors|`bool`|