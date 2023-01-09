# Table: github_repository_dependabot_secrets

The composite primary key for this table is (**org**, **repository_id**, **name**).

## Relations

This table depends on [github_repositories](github_repositories.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|repository_id (PK)|Int|
|name (PK)|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|visibility|String|
|selected_repositories_url|String|