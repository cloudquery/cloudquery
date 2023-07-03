# Table: gcp_websecurityscanner_scan_config_scan_run_crawled_urls

This table shows data for GCP Web Security Scanner Scan Config Scan Run Crawled Urls.

https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.crawledUrls/list#CrawledUrl

The composite primary key for this table is (**project_id**, **scan_run_name**, **http_method**, **url**).

## Relations

This table depends on [gcp_websecurityscanner_scan_config_scan_runs](gcp_websecurityscanner_scan_config_scan_runs).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|scan_run_name (PK)|`utf8`|
|http_method (PK)|`utf8`|
|url (PK)|`utf8`|
|body|`utf8`|