# Table: gcp_websecurityscanner_scan_config_scan_run_crawled_urls

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.crawledUrls/list#CrawledUrl

The composite primary key for this table is (**project_id**, **scan_run_name**, **http_method**, **url**).

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
|scan_run_name (PK)|String|
|http_method (PK)|String|
|url (PK)|String|
|body|String|