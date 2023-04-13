# Table: snyk_dependencies

This table shows data for Snyk Dependencies.

https://snyk.docs.apiary.io/#reference/dependencies/dependencies-by-organization/list-all-dependencies

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
|name|String|
|type|String|
|version|String|
|latest_version|String|
|latest_version_published_date|Timestamp|
|first_published_date|Timestamp|
|is_deprecated|Bool|
|deprecated_versions|StringArray|
|dependencies_with_issues|StringArray|
|issues_critical|Int|
|issues_high|Int|
|issues_medium|Int|
|issues_low|Int|
|licenses|JSON|
|projects|JSON|
|copyright|StringArray|