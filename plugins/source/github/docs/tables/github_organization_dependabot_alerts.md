# Table: github_organization_dependabot_alerts

The composite primary key for this table is (**org**, **number**).

## Relations

This table depends on [github_organizations](github_organizations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|number (PK)|Int|
|state|String|
|dependency|JSON|
|security_advisory|JSON|
|security_vulnerability|JSON|
|url|String|
|html_url|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|dismissed_at|Timestamp|
|dismissed_by|JSON|
|dismissed_reason|String|
|dismissed_comment|String|
|fixed_at|Timestamp|