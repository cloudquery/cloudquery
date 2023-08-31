# Table: github_repository_dependabot_alerts

This table shows data for Github Repository Dependabot Alerts.

The composite primary key for this table is (**org**, **repository_id**, **number**).

## Relations

This table depends on [github_repositories](github_repositories).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|number (PK)|`int64`|
|state|`utf8`|
|dependency|`json`|
|security_advisory|`json`|
|security_vulnerability|`json`|
|url|`utf8`|
|html_url|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|dismissed_at|`timestamp[us, tz=UTC]`|
|dismissed_by|`json`|
|dismissed_reason|`utf8`|
|dismissed_comment|`utf8`|
|fixed_at|`timestamp[us, tz=UTC]`|