# Table: github_organizations

This table shows data for Github Organizations.

The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_organizations:
  - [github_organization_dependabot_alerts](github_organization_dependabot_alerts)
  - [github_organization_dependabot_secrets](github_organization_dependabot_secrets)
  - [github_organization_members](github_organization_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|login|`utf8`|
|id (PK)|`int64`|
|node_id|`utf8`|
|avatar_url|`utf8`|
|html_url|`utf8`|
|name|`utf8`|
|company|`utf8`|
|blog|`utf8`|
|location|`utf8`|
|email|`utf8`|
|twitter_username|`utf8`|
|description|`utf8`|
|public_repos|`int64`|
|public_gists|`int64`|
|followers|`int64`|
|following|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|total_private_repos|`int64`|
|owned_private_repos|`int64`|
|private_gists|`int64`|
|disk_usage|`int64`|
|collaborators|`int64`|
|billing_email|`utf8`|
|type|`utf8`|
|plan|`json`|
|two_factor_requirement_enabled|`bool`|
|is_verified|`bool`|
|has_organization_projects|`bool`|
|has_repository_projects|`bool`|
|default_repository_permission|`utf8`|
|default_repository_settings|`utf8`|
|members_can_create_repositories|`bool`|
|members_can_create_public_repositories|`bool`|
|members_can_create_private_repositories|`bool`|
|members_can_create_internal_repositories|`bool`|
|members_can_fork_private_repositories|`bool`|
|members_allowed_repository_creation_type|`utf8`|
|members_can_create_pages|`bool`|
|members_can_create_public_pages|`bool`|
|members_can_create_private_pages|`bool`|
|web_commit_signoff_required|`bool`|
|advanced_security_enabled_for_new_repositories|`bool`|
|dependabot_alerts_enabled_for_new_repositories|`bool`|
|dependabot_security_updates_enabled_for_new_repositories|`bool`|
|dependency_graph_enabled_for_new_repositories|`bool`|
|secret_scanning_enabled_for_new_repositories|`bool`|
|secret_scanning_push_protection_enabled_for_new_repositories|`bool`|
|url|`utf8`|
|events_url|`utf8`|
|hooks_url|`utf8`|
|issues_url|`utf8`|
|members_url|`utf8`|
|public_members_url|`utf8`|
|repos_url|`utf8`|