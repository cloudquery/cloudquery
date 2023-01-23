# Table: snyk_dependencies

https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Dependency

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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