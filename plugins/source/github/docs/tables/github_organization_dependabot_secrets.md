# Table: github_organization_dependabot_secrets

The composite primary key for this table is (**org**, **name**).

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
|name (PK)|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|visibility|String|
|selected_repositories_url|String|