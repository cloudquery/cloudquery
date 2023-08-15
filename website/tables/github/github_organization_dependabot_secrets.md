# Table: github_organization_dependabot_secrets

This table shows data for Github Organization Dependabot Secrets.

The composite primary key for this table is (**org**, **name**).

## Relations

This table depends on [github_organizations](github_organizations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|name (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|visibility|`utf8`|
|selected_repositories_url|`utf8`|