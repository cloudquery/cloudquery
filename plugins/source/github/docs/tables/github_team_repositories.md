
# Table: github_team_repositories
Repository represents a GitHub repository.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|team_cq_id|uuid|Unique CloudQuery ID of github_teams table (FK)|
|team_id|bigint|The id of the team|
|id|bigint||
|node_id|text||
|owner_login|text||
|owner_id|bigint||
|owner_node_id|text||
|owner_avatar_url|text||
|owner_html_url|text||
|owner_gravatar_id|text||
|owner_name|text||
|owner_company|text||
|owner_blog|text||
|owner_location|text||
|owner_email|text||
|owner_hireable|boolean||
|owner_bio|text||
|owner_twitter_username|text||
|owner_public_repos|bigint||
|owner_public_gists|bigint||
|owner_followers|bigint||
|owner_following|bigint||
|owner_created_at_time|timestamp without time zone||
|owner_updated_at_time|timestamp without time zone||
|owner_suspended_at_time|timestamp without time zone||
|owner_type|text||
|owner_site_admin|boolean||
|owner_total_private_repos|bigint||
|owner_owned_private_repos|bigint||
|owner_private_gists|bigint||
|owner_disk_usage|bigint||
|owner_collaborators|bigint||
|owner_two_factor_authentication|boolean||
|owner_plan_name|text||
|owner_plan_space|bigint||
|owner_plan_collaborators|bigint||
|owner_plan_private_repos|bigint||
|owner_plan_filled_seats|bigint||
|owner_plan_seats|bigint||
|owner_ldap_dn|text||
|owner_url|text|API URLs|
|owner_events_url|text||
|owner_following_url|text||
|owner_followers_url|text||
|owner_gists_url|text||
|owner_organizations_url|text||
|owner_received_events_url|text||
|owner_repos_url|text||
|owner_starred_url|text||
|owner_subscriptions_url|text||
|owner_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|owner_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|owner_role_name|text||
|name|text||
|full_name|text||
|description|text||
|homepage|text||
|code_of_conduct_name|text||
|code_of_conduct_key|text||
|code_of_conduct_url|text||
|code_of_conduct_body|text||
|default_branch|text||
|master_branch|text||
|created_at_time|timestamp without time zone||
|pushed_at_time|timestamp without time zone||
|updated_at_time|timestamp without time zone||
|html_url|text||
|clone_url|text||
|git_url|text||
|mirror_url|text||
|ssh_url|text||
|s_v_n_url|text||
|language|text||
|fork|boolean||
|forks_count|bigint||
|network_count|bigint||
|open_issues_count|bigint||
|open_issues|bigint|Deprecated: Replaced by OpenIssuesCount|
|stargazers_count|bigint||
|subscribers_count|bigint||
|watchers_count|bigint|Deprecated: Replaced by StargazersCount|
|watchers|bigint|Deprecated: Replaced by StargazersCount|
|size|bigint||
|auto_init|boolean||
|parent|bigint||
|source|bigint||
|template_repository|bigint||
|organization_login|text||
|organization_id|bigint||
|organization_node_id|text||
|organization_avatar_url|text||
|organization_html_url|text||
|organization_name|text||
|organization_company|text||
|organization_blog|text||
|organization_location|text||
|organization_email|text||
|organization_twitter_username|text||
|organization_description|text||
|organization_public_repos|bigint||
|organization_public_gists|bigint||
|organization_followers|bigint||
|organization_following|bigint||
|organization_created_at|timestamp without time zone||
|organization_updated_at|timestamp without time zone||
|organization_total_private_repos|bigint||
|organization_owned_private_repos|bigint||
|organization_private_gists|bigint||
|organization_disk_usage|bigint||
|organization_collaborators|bigint||
|organization_billing_email|text||
|organization_type|text||
|organization_plan_name|text||
|organization_plan_space|bigint||
|organization_plan_collaborators|bigint||
|organization_plan_private_repos|bigint||
|organization_plan_filled_seats|bigint||
|organization_plan_seats|bigint||
|organization_two_factor_requirement_enabled|boolean||
|organization_is_verified|boolean||
|organization_has_organization_projects|boolean||
|organization_has_repository_projects|boolean||
|organization_default_repo_permission|text|DefaultRepoPermission can be one of: "read", "write", "admin", or "none"|
|organization_default_repo_settings|text|DefaultRepoSettings can be one of: "read", "write", "admin", or "none"|
|organization_members_can_create_repos|boolean|MembersCanCreateRepos default value is true and is only used in Organizations.Edit.|
|organization_members_can_create_public_repos|boolean|https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api|
|organization_members_can_create_private_repos|boolean||
|organization_members_can_create_internal_repos|boolean||
|organization_members_can_fork_private_repos|boolean|MembersCanForkPrivateRepos toggles whether organization members can fork private organization repositories.|
|organization_members_allowed_repository_creation_type|text|MembersAllowedRepositoryCreationType denotes if organization members can create repositories and the type of repositories they can create|
|organization_members_can_create_pages|boolean|MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.|
|organization_members_can_create_public_pages|boolean|MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.|
|organization_members_can_create_private_pages|boolean|MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.|
|organization_url|text|API URLs|
|organization_events_url|text||
|organization_hooks_url|text||
|organization_issues_url|text||
|organization_members_url|text||
|organization_public_members_url|text||
|organization_repos_url|text||
|permissions|jsonb||
|allow_rebase_merge|boolean||
|allow_update_branch|boolean||
|allow_squash_merge|boolean||
|allow_merge_commit|boolean||
|allow_auto_merge|boolean||
|allow_forking|boolean||
|delete_branch_on_merge|boolean||
|use_squash_p_r_title_as_default|boolean||
|topics|text[]||
|archived|boolean||
|disabled|boolean||
|license_key|text||
|license_name|text||
|license_url|text||
|license_s_p_d_x_id|text||
|license_html_url|text||
|license_featured|boolean||
|license_description|text||
|license_implementation|text||
|license_permissions|text[]||
|license_conditions|text[]||
|license_limitations|text[]||
|license_body|text||
|private|boolean|Additional mutable fields when creating and editing a repository|
|has_issues|boolean||
|has_wiki|boolean||
|has_pages|boolean||
|has_projects|boolean||
|has_downloads|boolean||
|is_template|boolean||
|license_template|text||
|gitignore_template|text||
|security_and_analysis_advanced_security_status|text||
|security_and_analysis_secret_scanning_status|text||
|url|text|API URLs|
|archive_url|text||
|assignees_url|text||
|blobs_url|text||
|branches_url|text||
|collaborators_url|text||
|comments_url|text||
|commits_url|text||
|compare_url|text||
|contents_url|text||
|contributors_url|text||
|deployments_url|text||
|downloads_url|text||
|events_url|text||
|forks_url|text||
|git_commits_url|text||
|git_refs_url|text||
|git_tags_url|text||
|hooks_url|text||
|issue_comment_url|text||
|issue_events_url|text||
|issues_url|text||
|keys_url|text||
|labels_url|text||
|languages_url|text||
|merges_url|text||
|milestones_url|text||
|notifications_url|text||
|pulls_url|text||
|releases_url|text||
|stargazers_url|text||
|statuses_url|text||
|subscribers_url|text||
|subscription_url|text||
|tags_url|text||
|trees_url|text||
|teams_url|text||
|text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|visibility|text|Visibility is only used for Create and Edit endpoints|
|role_name|text|RoleName is only returned by the API 'check team permissions for a repository'. See: teams.go (IsTeamRepoByID) https://docs.github.com/en/rest/teams/teams#check-team-permissions-for-a-repository|
