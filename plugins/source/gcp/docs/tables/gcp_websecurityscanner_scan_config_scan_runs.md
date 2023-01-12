# Table: gcp_websecurityscanner_scan_config_scan_runs

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_websecurityscanner_scan_configs](gcp_websecurityscanner_scan_configs.md).

The following tables depend on gcp_websecurityscanner_scan_config_scan_runs:
  - [gcp_websecurityscanner_scan_config_scan_run_crawled_urls](gcp_websecurityscanner_scan_config_scan_run_crawled_urls.md)
  - [gcp_websecurityscanner_scan_config_scan_run_findings](gcp_websecurityscanner_scan_config_scan_run_findings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|execution_state|String|
|result_state|String|
|start_time|Timestamp|
|end_time|Timestamp|
|urls_crawled_count|Int|
|urls_tested_count|Int|
|has_vulnerabilities|Bool|
|progress_percent|Int|
|error_trace|JSON|
|warning_traces|JSON|