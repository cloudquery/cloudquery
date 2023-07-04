# Table: gcp_websecurityscanner_scan_config_scan_runs

This table shows data for GCP Web Security Scanner Scan Config Scan Runs.

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_websecurityscanner_scan_configs](gcp_websecurityscanner_scan_configs).

The following tables depend on gcp_websecurityscanner_scan_config_scan_runs:
  - [gcp_websecurityscanner_scan_config_scan_run_crawled_urls](gcp_websecurityscanner_scan_config_scan_run_crawled_urls)
  - [gcp_websecurityscanner_scan_config_scan_run_findings](gcp_websecurityscanner_scan_config_scan_run_findings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|execution_state|`utf8`|
|result_state|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|urls_crawled_count|`int64`|
|urls_tested_count|`int64`|
|has_vulnerabilities|`bool`|
|progress_percent|`int64`|
|error_trace|`json`|
|warning_traces|`json`|