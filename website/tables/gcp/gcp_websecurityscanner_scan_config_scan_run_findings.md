# Table: gcp_websecurityscanner_scan_config_scan_run_findings

This table shows data for GCP Web Security Scanner Scan Config Scan Run Findings.

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.findings

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_websecurityscanner_scan_config_scan_runs](gcp_websecurityscanner_scan_config_scan_runs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|finding_type|`utf8`|
|severity|`utf8`|
|http_method|`utf8`|
|fuzzed_url|`utf8`|
|body|`utf8`|
|description|`utf8`|
|reproduction_url|`utf8`|
|frame_url|`utf8`|
|final_url|`utf8`|
|tracking_id|`utf8`|
|form|`json`|
|outdated_library|`json`|
|violating_resource|`json`|
|vulnerable_headers|`json`|
|vulnerable_parameters|`json`|
|xss|`json`|
|xxe|`json`|