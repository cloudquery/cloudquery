# Table: snyk_reporting_issues

This table shows data for Snyk Reporting Issues.

https://snyk.docs.apiary.io/#reference/reporting-api/get-list-of-latest-issues

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|organization_id (PK)|String|
|id (PK)|String|
|issue|JSON|
|projects|JSON|
|project|JSON|
|is_fixed|Bool|
|introduced_date|String|
|patched_date|String|
|fixed_date|String|