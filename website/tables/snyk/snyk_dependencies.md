# Table: snyk_dependencies

This table shows data for Snyk Dependencies.

https://snyk.docs.apiary.io/#reference/dependencies/dependencies-by-organization/list-all-dependencies

The composite primary key for this table is (**organization_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|organization_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|
|version|`utf8`|
|latest_version|`utf8`|
|latest_version_published_date|`timestamp[us, tz=UTC]`|
|first_published_date|`timestamp[us, tz=UTC]`|
|is_deprecated|`bool`|
|deprecated_versions|`list<item: utf8, nullable>`|
|dependencies_with_issues|`list<item: utf8, nullable>`|
|issues_critical|`int64`|
|issues_high|`int64`|
|issues_medium|`int64`|
|issues_low|`int64`|
|licenses|`json`|
|projects|`json`|
|copyright|`list<item: utf8, nullable>`|