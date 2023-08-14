# Table: github_organization_dependabot_alerts

This table shows data for Github Organization Dependabot Alerts.

The composite primary key for this table is (**org**, **html_url**).

## Relations

This table depends on [github_organizations](github_organizations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|number|`int64`|
|state|`utf8`|
|dependency|`json`|
|security_advisory|`json`|
|security_vulnerability|`json`|
|url|`utf8`|
|html_url (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|dismissed_at|`timestamp[us, tz=UTC]`|
|dismissed_by|`json`|
|dismissed_reason|`utf8`|
|dismissed_comment|`utf8`|
|fixed_at|`timestamp[us, tz=UTC]`|