# Table: github_repository_dependabot_secrets

This table shows data for Github Repository Dependabot Secrets.

The composite primary key for this table is (**org**, **repository_id**, **name**).

## Relations

This table depends on [github_repositories](github_repositories).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|name (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|visibility|`utf8`|
|selected_repositories_url|`utf8`|