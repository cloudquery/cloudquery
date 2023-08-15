# Table: snyk_reporting_issues

This table shows data for Snyk Reporting Issues.

https://snyk.docs.apiary.io/#reference/reporting-api/issues/get-list-of-issues

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|id (PK)|`utf8`|
|issue|`json`|
|projects|`json`|
|project|`json`|
|is_fixed|`bool`|
|introduced_date|`utf8`|
|patched_date|`utf8`|
|fixed_date|`utf8`|