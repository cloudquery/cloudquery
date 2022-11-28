# Table: github_organizations



The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_organizations:
  - [github_organization_members](github_organization_members.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|id (PK)|Int|
|login|String|
|node_id|String|
|avatar_url|String|
|html_url|String|
|name|String|
|company|String|
|blog|String|
|location|String|
|email|String|
|twitter_username|String|
|description|String|
|public_repos|Int|
|public_gists|Int|
|followers|Int|
|following|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|total_private_repos|Int|
|owned_private_repos|Int|
|private_gists|Int|
|disk_usage|Int|
|collaborators|Int|
|billing_email|String|
|type|String|
|plan|JSON|
|two_factor_requirement_enabled|Bool|
|is_verified|Bool|
|has_organization_projects|Bool|
|has_repository_projects|Bool|
|default_repository_permission|String|
|default_repository_settings|String|
|members_can_create_repositories|Bool|
|members_can_create_public_repositories|Bool|
|members_can_create_private_repositories|Bool|
|members_can_create_internal_repositories|Bool|
|members_can_fork_private_repositories|Bool|
|members_allowed_repository_creation_type|String|
|members_can_create_pages|Bool|
|members_can_create_public_pages|Bool|
|members_can_create_private_pages|Bool|
|url|String|
|events_url|String|
|hooks_url|String|
|issues_url|String|
|members_url|String|
|public_members_url|String|
|repos_url|String|