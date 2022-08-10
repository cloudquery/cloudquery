
# Table: github_installations
Installation represents a GitHub Apps installation.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|org|text|The Github Organization of the resource.|
|id|bigint||
|node_id|text||
|app_id|bigint||
|app_slug|text||
|target_id|bigint||
|account_login|text||
|account_id|bigint||
|account_node_id|text||
|account_avatar_url|text||
|account_html_url|text||
|account_gravatar_id|text||
|account_name|text||
|account_company|text||
|account_blog|text||
|account_location|text||
|account_email|text||
|account_hireable|boolean||
|account_bio|text||
|account_twitter_username|text||
|account_public_repos|bigint||
|account_public_gists|bigint||
|account_followers|bigint||
|account_following|bigint||
|account_created_at_time|timestamp without time zone||
|account_updated_at_time|timestamp without time zone||
|account_suspended_at_time|timestamp without time zone||
|account_type|text||
|account_site_admin|boolean||
|account_total_private_repos|bigint||
|account_owned_private_repos|bigint||
|account_private_gists|bigint||
|account_disk_usage|bigint||
|account_collaborators|bigint||
|account_two_factor_authentication|boolean||
|account_plan_name|text||
|account_plan_space|bigint||
|account_plan_collaborators|bigint||
|account_plan_private_repos|bigint||
|account_plan_filled_seats|bigint||
|account_plan_seats|bigint||
|account_ldap_dn|text||
|account_url|text|API URLs|
|account_events_url|text||
|account_following_url|text||
|account_followers_url|text||
|account_gists_url|text||
|account_organizations_url|text||
|account_received_events_url|text||
|account_repos_url|text||
|account_starred_url|text||
|account_subscriptions_url|text||
|account_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|account_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|account_role_name|text||
|access_tokens_url|text||
|repositories_url|text||
|html_url|text||
|target_type|text||
|single_file_name|text||
|repository_selection|text||
|events|text[]||
|single_file_paths|text[]||
|permissions_actions|text||
|permissions_administration|text||
|permissions_blocking|text||
|permissions_checks|text||
|permissions_contents|text||
|permissions_content_references|text||
|permissions_deployments|text||
|permissions_emails|text||
|permissions_environments|text||
|permissions_followers|text||
|permissions_issues|text||
|permissions_metadata|text||
|permissions_members|text||
|permissions_organization_administration|text||
|permissions_organization_hooks|text||
|permissions_organization_plan|text||
|permissions_organization_pre_receive_hooks|text||
|permissions_organization_projects|text||
|permissions_organization_secrets|text||
|permissions_organization_self_hosted_runners|text||
|permissions_organization_user_blocking|text||
|permissions_packages|text||
|permissions_pages|text||
|permissions_pull_requests|text||
|permissions_repository_hooks|text||
|permissions_repository_projects|text||
|permissions_repository_pre_receive_hooks|text||
|permissions_secrets|text||
|permissions_secret_scanning_alerts|text||
|permissions_security_events|text||
|permissions_single_file|text||
|permissions_statuses|text||
|permissions_team_discussions|text||
|permissions_vulnerability_alerts|text||
|permissions_workflows|text||
|created_at_time|timestamp without time zone||
|updated_at_time|timestamp without time zone||
|has_multiple_single_files|boolean||
|suspended_by_login|text||
|suspended_by_id|bigint||
|suspended_by_node_id|text||
|suspended_by_avatar_url|text||
|suspended_by_html_url|text||
|suspended_by_gravatar_id|text||
|suspended_by_name|text||
|suspended_by_company|text||
|suspended_by_blog|text||
|suspended_by_location|text||
|suspended_by_email|text||
|suspended_by_hireable|boolean||
|suspended_by_bio|text||
|suspended_by_twitter_username|text||
|suspended_by_public_repos|bigint||
|suspended_by_public_gists|bigint||
|suspended_by_followers|bigint||
|suspended_by_following|bigint||
|suspended_by_created_at_time|timestamp without time zone||
|suspended_by_updated_at_time|timestamp without time zone||
|suspended_by_suspended_at_time|timestamp without time zone||
|suspended_by_type|text||
|suspended_by_site_admin|boolean||
|suspended_by_total_private_repos|bigint||
|suspended_by_owned_private_repos|bigint||
|suspended_by_private_gists|bigint||
|suspended_by_disk_usage|bigint||
|suspended_by_collaborators|bigint||
|suspended_by_two_factor_authentication|boolean||
|suspended_by_plan_name|text||
|suspended_by_plan_space|bigint||
|suspended_by_plan_collaborators|bigint||
|suspended_by_plan_private_repos|bigint||
|suspended_by_plan_filled_seats|bigint||
|suspended_by_plan_seats|bigint||
|suspended_by_ldap_dn|text||
|suspended_by_url|text|API URLs|
|suspended_by_events_url|text||
|suspended_by_following_url|text||
|suspended_by_followers_url|text||
|suspended_by_gists_url|text||
|suspended_by_organizations_url|text||
|suspended_by_received_events_url|text||
|suspended_by_repos_url|text||
|suspended_by_starred_url|text||
|suspended_by_subscriptions_url|text||
|suspended_by_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|suspended_by_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|suspended_by_role_name|text||
|suspended_at_time|timestamp without time zone||
