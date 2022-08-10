
# Table: github_organizations
Organization represents a GitHub organization account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|login|text||
|id|bigint||
|node_id|text||
|avatar_url|text||
|html_url|text||
|name|text||
|company|text||
|blog|text||
|location|text||
|email|text||
|twitter_username|text||
|description|text||
|public_repos|bigint||
|public_gists|bigint||
|followers|bigint||
|following|bigint||
|created_at|timestamp without time zone||
|updated_at|timestamp without time zone||
|total_private_repos|bigint||
|owned_private_repos|bigint||
|private_gists|bigint||
|disk_usage|bigint||
|collaborators|bigint||
|billing_email|text||
|type|text||
|plan_name|text||
|plan_space|bigint||
|plan_collaborators|bigint||
|plan_private_repos|bigint||
|plan_filled_seats|bigint||
|plan_seats|bigint||
|two_factor_requirement_enabled|boolean||
|is_verified|boolean||
|has_organization_projects|boolean||
|has_repository_projects|boolean||
|default_repo_permission|text|DefaultRepoPermission can be one of: "read", "write", "admin", or "none"|
|default_repo_settings|text|DefaultRepoSettings can be one of: "read", "write", "admin", or "none"|
|members_can_create_repos|boolean|MembersCanCreateRepos default value is true and is only used in Organizations.Edit.|
|members_can_create_public_repos|boolean|https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api|
|members_can_create_private_repos|boolean||
|members_can_create_internal_repos|boolean||
|members_can_fork_private_repos|boolean|MembersCanForkPrivateRepos toggles whether organization members can fork private organization repositories.|
|members_allowed_repository_creation_type|text|MembersAllowedRepositoryCreationType denotes if organization members can create repositories and the type of repositories they can create|
|members_can_create_pages|boolean|MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.|
|members_can_create_public_pages|boolean|MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.|
|members_can_create_private_pages|boolean|MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.|
|url|text|API URLs|
|events_url|text||
|hooks_url|text||
|issues_url|text||
|members_url|text||
|public_members_url|text||
|repos_url|text||
