
# Table: github_issues
Issue represents a GitHub issue on a repository.  Note: As far as the GitHub API is concerned, every pull request is an issue, but not every issue is a pull request
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|org|text|The Github Organization of the resource.|
|id|bigint||
|number|bigint||
|state|text||
|locked|boolean||
|title|text||
|body|text||
|author_association|text||
|user_login|text||
|user_id|bigint||
|user_node_id|text||
|user_avatar_url|text||
|user_html_url|text||
|user_gravatar_id|text||
|user_name|text||
|user_company|text||
|user_blog|text||
|user_location|text||
|user_email|text||
|user_hireable|boolean||
|user_bio|text||
|user_twitter_username|text||
|user_public_repos|bigint||
|user_public_gists|bigint||
|user_followers|bigint||
|user_following|bigint||
|user_created_at_time|timestamp without time zone||
|user_updated_at_time|timestamp without time zone||
|user_suspended_at_time|timestamp without time zone||
|user_type|text||
|user_site_admin|boolean||
|user_total_private_repos|bigint||
|user_owned_private_repos|bigint||
|user_private_gists|bigint||
|user_disk_usage|bigint||
|user_collaborators|bigint||
|user_two_factor_authentication|boolean||
|user_plan_name|text||
|user_plan_space|bigint||
|user_plan_collaborators|bigint||
|user_plan_private_repos|bigint||
|user_plan_filled_seats|bigint||
|user_plan_seats|bigint||
|user_ldap_dn|text||
|user_url|text|API URLs|
|user_events_url|text||
|user_following_url|text||
|user_followers_url|text||
|user_gists_url|text||
|user_organizations_url|text||
|user_received_events_url|text||
|user_repos_url|text||
|user_starred_url|text||
|user_subscriptions_url|text||
|user_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|user_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|user_role_name|text||
|assignee_login|text||
|assignee_id|bigint||
|assignee_node_id|text||
|assignee_avatar_url|text||
|assignee_html_url|text||
|assignee_gravatar_id|text||
|assignee_name|text||
|assignee_company|text||
|assignee_blog|text||
|assignee_location|text||
|assignee_email|text||
|assignee_hireable|boolean||
|assignee_bio|text||
|assignee_twitter_username|text||
|assignee_public_repos|bigint||
|assignee_public_gists|bigint||
|assignee_followers|bigint||
|assignee_following|bigint||
|assignee_created_at_time|timestamp without time zone||
|assignee_updated_at_time|timestamp without time zone||
|assignee_suspended_at_time|timestamp without time zone||
|assignee_type|text||
|assignee_site_admin|boolean||
|assignee_total_private_repos|bigint||
|assignee_owned_private_repos|bigint||
|assignee_private_gists|bigint||
|assignee_disk_usage|bigint||
|assignee_collaborators|bigint||
|assignee_two_factor_authentication|boolean||
|assignee_plan_name|text||
|assignee_plan_space|bigint||
|assignee_plan_collaborators|bigint||
|assignee_plan_private_repos|bigint||
|assignee_plan_filled_seats|bigint||
|assignee_plan_seats|bigint||
|assignee_ldap_dn|text||
|assignee_url|text|API URLs|
|assignee_events_url|text||
|assignee_following_url|text||
|assignee_followers_url|text||
|assignee_gists_url|text||
|assignee_organizations_url|text||
|assignee_received_events_url|text||
|assignee_repos_url|text||
|assignee_starred_url|text||
|assignee_subscriptions_url|text||
|assignee_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|assignee_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|assignee_role_name|text||
|comments|bigint||
|closed_at|timestamp without time zone||
|created_at|timestamp without time zone||
|updated_at|timestamp without time zone||
|closed_by_login|text||
|closed_by_id|bigint||
|closed_by_node_id|text||
|closed_by_avatar_url|text||
|closed_by_html_url|text||
|closed_by_gravatar_id|text||
|closed_by_name|text||
|closed_by_company|text||
|closed_by_blog|text||
|closed_by_location|text||
|closed_by_email|text||
|closed_by_hireable|boolean||
|closed_by_bio|text||
|closed_by_twitter_username|text||
|closed_by_public_repos|bigint||
|closed_by_public_gists|bigint||
|closed_by_followers|bigint||
|closed_by_following|bigint||
|closed_by_created_at_time|timestamp without time zone||
|closed_by_updated_at_time|timestamp without time zone||
|closed_by_suspended_at_time|timestamp without time zone||
|closed_by_type|text||
|closed_by_site_admin|boolean||
|closed_by_total_private_repos|bigint||
|closed_by_owned_private_repos|bigint||
|closed_by_private_gists|bigint||
|closed_by_disk_usage|bigint||
|closed_by_collaborators|bigint||
|closed_by_two_factor_authentication|boolean||
|closed_by_plan_name|text||
|closed_by_plan_space|bigint||
|closed_by_plan_collaborators|bigint||
|closed_by_plan_private_repos|bigint||
|closed_by_plan_filled_seats|bigint||
|closed_by_plan_seats|bigint||
|closed_by_ldap_dn|text||
|closed_by_url|text|API URLs|
|closed_by_events_url|text||
|closed_by_following_url|text||
|closed_by_followers_url|text||
|closed_by_gists_url|text||
|closed_by_organizations_url|text||
|closed_by_received_events_url|text||
|closed_by_repos_url|text||
|closed_by_starred_url|text||
|closed_by_subscriptions_url|text||
|closed_by_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|closed_by_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|closed_by_role_name|text||
|url|text||
|html_url|text||
|comments_url|text||
|events_url|text||
|labels_url|text||
|repository_url|text||
|milestone_url|text||
|milestone_html_url|text||
|milestone_labels_url|text||
|milestone_id|bigint||
|milestone_number|bigint||
|milestone_state|text||
|milestone_title|text||
|milestone_description|text||
|milestone_creator_login|text||
|milestone_creator_id|bigint||
|milestone_creator_node_id|text||
|milestone_creator_avatar_url|text||
|milestone_creator_html_url|text||
|milestone_creator_gravatar_id|text||
|milestone_creator_name|text||
|milestone_creator_company|text||
|milestone_creator_blog|text||
|milestone_creator_location|text||
|milestone_creator_email|text||
|milestone_creator_hireable|boolean||
|milestone_creator_bio|text||
|milestone_creator_twitter_username|text||
|milestone_creator_public_repos|bigint||
|milestone_creator_public_gists|bigint||
|milestone_creator_followers|bigint||
|milestone_creator_following|bigint||
|milestone_creator_created_at_time|timestamp without time zone||
|milestone_creator_updated_at_time|timestamp without time zone||
|milestone_creator_suspended_at_time|timestamp without time zone||
|milestone_creator_type|text||
|milestone_creator_site_admin|boolean||
|milestone_creator_total_private_repos|bigint||
|milestone_creator_owned_private_repos|bigint||
|milestone_creator_private_gists|bigint||
|milestone_creator_disk_usage|bigint||
|milestone_creator_collaborators|bigint||
|milestone_creator_two_factor_authentication|boolean||
|milestone_creator_plan_name|text||
|milestone_creator_plan_space|bigint||
|milestone_creator_plan_collaborators|bigint||
|milestone_creator_plan_private_repos|bigint||
|milestone_creator_plan_filled_seats|bigint||
|milestone_creator_plan_seats|bigint||
|milestone_creator_ldap_dn|text||
|milestone_creator_url|text|API URLs|
|milestone_creator_events_url|text||
|milestone_creator_following_url|text||
|milestone_creator_followers_url|text||
|milestone_creator_gists_url|text||
|milestone_creator_organizations_url|text||
|milestone_creator_received_events_url|text||
|milestone_creator_repos_url|text||
|milestone_creator_starred_url|text||
|milestone_creator_subscriptions_url|text||
|milestone_creator_text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|milestone_creator_permissions|jsonb|Permissions and RoleName identify the permissions and role that a user has on a given repository|
|milestone_creator_role_name|text||
|milestone_open_issues|bigint||
|milestone_closed_issues|bigint||
|milestone_created_at|timestamp without time zone||
|milestone_updated_at|timestamp without time zone||
|milestone_closed_at|timestamp without time zone||
|milestone_due_on|timestamp without time zone||
|milestone_node_id|text||
|pull_request_links_url|text||
|pull_request_links_html_url|text||
|pull_request_links_diff_url|text||
|pull_request_links_patch_url|text||
|repository_id|bigint||
|reactions_total_count|bigint||
|reactions_plus_one|bigint||
|reactions_laugh|bigint||
|reactions_confused|bigint||
|reactions_heart|bigint||
|reactions_hooray|bigint||
|reactions_rocket|bigint||
|reactions_eyes|bigint||
|reactions_url|text||
|node_id|text||
|text_matches|jsonb|TextMatches is only populated from search results that request text matches See: search.go and https://docs.github.com/en/rest/search/#text-match-metadata|
|active_lock_reason|text|ActiveLockReason is populated only when LockReason is provided while locking the issue. Possible values are: "off-topic", "too heated", "resolved", and "spam".|
