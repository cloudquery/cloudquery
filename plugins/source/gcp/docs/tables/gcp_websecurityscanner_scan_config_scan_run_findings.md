# Table: gcp_websecurityscanner_scan_config_scan_run_findings

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.findings

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_websecurityscanner_scan_config_scan_runs](gcp_websecurityscanner_scan_config_scan_runs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|finding_type|String|
|severity|String|
|http_method|String|
|fuzzed_url|String|
|body|String|
|description|String|
|reproduction_url|String|
|frame_url|String|
|final_url|String|
|tracking_id|String|
|form|JSON|
|outdated_library|JSON|
|violating_resource|JSON|
|vulnerable_headers|JSON|
|vulnerable_parameters|JSON|
|xss|JSON|
|xxe|JSON|